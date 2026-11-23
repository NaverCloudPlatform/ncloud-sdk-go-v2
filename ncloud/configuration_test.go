package ncloud

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud/credentials"
)

func TestGetCredentialsRecoversConfiguredProvider(t *testing.T) {
	var ready int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&ready) == 0 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		fmt.Fprintf(w, `{"accessKey":"AK","secretKey":"SK","expiration":%q}`,
			time.Now().Add(time.Hour).Format(time.RFC3339))
	}))
	defer srv.Close()

	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("token-value"), 0600); err != nil {
		t.Fatal(err)
	}
	p := &credentials.PodIdentityProvider{TokenFile: tokenFile, APIEndpoint: srv.URL}
	cfg := &Configuration{Credentials: credentials.LoadCredentials([]credentials.Provider{p})}

	if got := cfg.GetCredentials(); got != nil {
		t.Fatalf("Expected: nil, Actual: %v", got)
	}

	atomic.StoreInt32(&ready, 1)
	if got := cfg.GetCredentials(); got == nil || got.AccessKey() != "AK" {
		t.Fatalf("Expected: AK, Actual: %v", got)
	}
}

func TestInitCredentialsKeepsFailedProvider(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("token-value"), 0600); err != nil {
		t.Fatal(err)
	}
	p := &credentials.PodIdentityProvider{TokenFile: tokenFile, APIEndpoint: srv.URL}
	cfg := &Configuration{Credentials: credentials.LoadCredentials([]credentials.Provider{p})}

	cfg.InitCredentials()
	if got := cfg.Credentials.ProviderName(); got != p.Name() {
		t.Fatalf("Expected: %s, Actual: %s", p.Name(), got)
	}
}

func TestInitCredentialsLoadsWithoutProvider(t *testing.T) {
	setenv(t, "NCLOUD_ACCESS_KEY_ID", "AK")
	setenv(t, "NCLOUD_SECRET_ACCESS_KEY", "SK")
	cfg := &Configuration{Credentials: &credentials.Credentials{}}

	cfg.InitCredentials()
	if !cfg.ValidCredentials() || cfg.Credentials.AccessKey() != "AK" {
		t.Fatalf("Expected: AK, Actual: %s from %s", cfg.Credentials.AccessKey(), cfg.Credentials.ProviderName())
	}
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
	os.Setenv(key, value)
}
