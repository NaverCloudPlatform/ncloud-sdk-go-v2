package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	podIdentityTokenFileEnv   = "NCLOUD_POD_IDENTITY_TOKEN_FILE"
	podIdentityAPIEndpointEnv = "NCLOUD_POD_IDENTITY_API_ENDPOINT"

	podIdentityTimeout = 3 * time.Second
)

type PodIdentityProvider struct {
	// TokenFile and APIEndpoint default to NCLOUD_POD_IDENTITY_TOKEN_FILE
	// and NCLOUD_POD_IDENTITY_API_ENDPOINT, injected by the pod identity webhook.
	TokenFile   string
	APIEndpoint string
	HTTPClient  *http.Client
}

func (p *PodIdentityProvider) Name() string {
	return "PodIdentityProvider"
}

func (p *PodIdentityProvider) Retrieve() (Value, error) {
	tokenFile := p.TokenFile
	if tokenFile == "" {
		tokenFile = os.Getenv(podIdentityTokenFileEnv)
	}

	endpoint := p.APIEndpoint
	if endpoint == "" {
		endpoint = os.Getenv(podIdentityAPIEndpointEnv)
	}

	if tokenFile == "" {
		return Value{}, errors.New(podIdentityTokenFileEnv + " not found in environment")
	}
	if endpoint == "" {
		return Value{}, errors.New(podIdentityAPIEndpointEnv + " not found in environment")
	}

	// kubelet rotates the token file, so read it on every call.
	token, err := readPodIdentityToken(tokenFile)
	if err != nil {
		return Value{}, err
	}

	return p.exchange(endpoint, token)
}

func readPodIdentityToken(tokenFile string) (string, error) {
	token, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", fmt.Errorf("failed to read pod identity token %s: %v", tokenFile, err)
	}

	trimmed := strings.TrimSpace(string(token))
	if trimmed == "" {
		return "", fmt.Errorf("pod identity token %s is empty", tokenFile)
	}
	return trimmed, nil
}

type podIdentityCredentials struct {
	AccessKey  string    `json:"accessKey"`
	SecretKey  string    `json:"secretKey"`
	Expiration time.Time `json:"expiration"`
}

type podIdentityError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func (p *PodIdentityProvider) exchange(endpoint, token string) (Value, error) {
	client := p.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: podIdentityTimeout}
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return Value{}, fmt.Errorf("failed to build pod identity request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Value{}, fmt.Errorf("pod identity exchange request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		detail, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		var e podIdentityError
		if json.Unmarshal(detail, &e) == nil && e.Error != "" {
			return Value{}, fmt.Errorf("pod identity exchange returned %s: %s: %s",
				resp.Status, e.Error, e.Message)
		}
		return Value{}, fmt.Errorf("pod identity exchange returned %s: %s",
			resp.Status, strings.TrimSpace(string(detail)))
	}

	var creds podIdentityCredentials
	if err := json.NewDecoder(resp.Body).Decode(&creds); err != nil {
		return Value{}, fmt.Errorf("failed to decode pod identity credentials: %v", err)
	}

	if creds.AccessKey == "" || creds.SecretKey == "" {
		return Value{}, errors.New("pod identity exchange returned empty credentials")
	}
	if creds.Expiration.IsZero() {
		return Value{}, errors.New("pod identity exchange returned no expiration")
	}

	return Value{
		AccessKey:  creds.AccessKey,
		SecretKey:  creds.SecretKey,
		Expiration: creds.Expiration,
	}, nil
}
