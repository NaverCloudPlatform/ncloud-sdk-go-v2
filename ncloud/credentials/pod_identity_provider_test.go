package credentials

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeToken(t *testing.T, token string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(path, []byte(token), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

// setenv restores the previous value after the test. t.Setenv needs Go 1.17.
func setenv(t *testing.T, key, value string) {
	t.Helper()
	old, ok := os.LookupEnv(key)
	t.Cleanup(func() {
		if ok {
			os.Setenv(key, old)
		} else {
			os.Unsetenv(key)
		}
	})
	if value == "" {
		os.Unsetenv(key)
	} else {
		os.Setenv(key, value)
	}
}

func bodyServer(t *testing.T, status int, body string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func credentialsServer(t *testing.T, seen *[]string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if seen != nil {
			*seen = append(*seen, r.Header.Get("Authorization"))
		}
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"accessKey":"AK","secretKey":"SK","expiration":%q}`,
			time.Now().Add(time.Hour).Format(time.RFC3339))
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestPodIdentityProviderRetrieve(t *testing.T) {
	var seen []string
	srv := credentialsServer(t, &seen)

	p := &PodIdentityProvider{
		TokenFile:   writeToken(t, "token-value\n"),
		APIEndpoint: srv.URL,
	}

	value, err := p.Retrieve()
	if err != nil {
		t.Fatalf("Expected: nil, Actual: %v", err)
	}
	if value.AccessKey != "AK" || value.SecretKey != "SK" {
		t.Fatalf("Expected: AK/SK, Actual: %s/%s", value.AccessKey, value.SecretKey)
	}
	if left := time.Until(value.Expiration); left > time.Hour || left < time.Hour-time.Minute {
		t.Fatalf("Expected: expiration about %v ahead, Actual: %v", time.Hour, left)
	}
	if len(seen) != 1 || seen[0] != "Bearer token-value" {
		t.Fatalf("Expected: [Bearer token-value], Actual: %v", seen)
	}
}

func TestPodIdentityProviderRereadsTokenFile(t *testing.T) {
	var seen []string
	srv := credentialsServer(t, &seen)

	path := writeToken(t, "first")
	p := &PodIdentityProvider{TokenFile: path, APIEndpoint: srv.URL}

	if _, err := p.Retrieve(); err != nil {
		t.Fatalf("Expected: nil, Actual: %v", err)
	}
	if err := os.WriteFile(path, []byte("second"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Retrieve(); err != nil {
		t.Fatalf("Expected: nil, Actual: %v", err)
	}

	if len(seen) != 2 || seen[0] != "Bearer first" || seen[1] != "Bearer second" {
		t.Fatalf("Expected: [Bearer first Bearer second], Actual: %v", seen)
	}
}

func TestPodIdentityProviderReadsEnvironment(t *testing.T) {
	srv := credentialsServer(t, nil)
	setenv(t, podIdentityTokenFileEnv, writeToken(t, "token-value"))
	setenv(t, podIdentityAPIEndpointEnv, srv.URL)

	p := &PodIdentityProvider{}
	if _, err := p.Retrieve(); err != nil {
		t.Fatalf("Expected: nil, Actual: %v", err)
	}
}

func TestPodIdentityProviderErrors(t *testing.T) {
	setenv(t, podIdentityTokenFileEnv, "")
	setenv(t, podIdentityAPIEndpointEnv, "")

	ok := credentialsServer(t, nil)
	empty := bodyServer(t, http.StatusOK, `{"accessKey":"","secretKey":""}`)
	noExpiration := bodyServer(t, http.StatusOK, `{"accessKey":"AK","secretKey":"SK"}`)
	malformed := bodyServer(t, http.StatusOK, `not json`)

	tests := []struct {
		name      string
		tokenFile string
		endpoint  string
	}{
		{"missing token file setting", "", ok.URL},
		{"missing endpoint setting", writeToken(t, "token-value"), ""},
		{"missing token file", filepath.Join(t.TempDir(), "absent"), ok.URL},
		{"empty token file", writeToken(t, "  \n"), ok.URL},
		{"invalid endpoint", writeToken(t, "token-value"), "://invalid"},
		{"empty credentials", writeToken(t, "token-value"), empty.URL},
		{"missing expiration", writeToken(t, "token-value"), noExpiration.URL},
		{"malformed response", writeToken(t, "token-value"), malformed.URL},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PodIdentityProvider{TokenFile: tt.tokenFile, APIEndpoint: tt.endpoint}
			if _, err := p.Retrieve(); err == nil {
				t.Fatal("Expected: error, Actual: nil")
			}
		})
	}
}

// statusServer answers with the given statuses in order, then succeeds.
func statusServer(t *testing.T, seen *[]string, statuses ...int) *httptest.Server {
	t.Helper()
	ok := credentialsServer(t, nil)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*seen = append(*seen, r.Header.Get("Authorization"))
		if len(*seen) <= len(statuses) {
			w.WriteHeader(statuses[len(*seen)-1])
			fmt.Fprint(w, `{"error":"Rejected","message":"rejected"}`)
			return
		}
		ok.Config.Handler.ServeHTTP(w, r)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestPodIdentityProviderServerErrors(t *testing.T) {
	for _, status := range []int{400, 401, 403, 429, 500, 503} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var seen []string
			srv := statusServer(t, &seen, status)
			p := &PodIdentityProvider{TokenFile: writeToken(t, "token-value"), APIEndpoint: srv.URL}

			if _, err := p.Retrieve(); err == nil {
				t.Fatal("Expected: error, Actual: nil")
			}
			if len(seen) != 1 {
				t.Fatalf("Expected: 1 request, Actual: %d", len(seen))
			}
		})
	}
}

func TestPodIdentityProviderErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{"exchange server error", `{"error":"NoAssociation","message":"no role"}`, "NoAssociation: no role"},
		{"apigw error", `{"error":{"errorCode":"200","message":"Authentication Failed"}}`, `"errorCode":"200"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := bodyServer(t, http.StatusForbidden, tt.body)
			p := &PodIdentityProvider{TokenFile: writeToken(t, "token-value"), APIEndpoint: srv.URL}

			_, err := p.Retrieve()
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("Expected: error containing %s, Actual: %v", tt.want, err)
			}
		})
	}
}
