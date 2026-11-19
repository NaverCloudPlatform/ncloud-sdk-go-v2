package credentials

import (
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud/metadata"
	"sync"
	"time"
)

type Credentials struct {
	mu          sync.RWMutex
	value       Value
	provider    Provider
	lastRefresh time.Time
	lastErr     error
}

type Value struct {
	AccessKey  string
	SecretKey  string
	Expiration time.Time
}

// IsExpired returns if the credentials are expired.
func (c *Credentials) IsExpired() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.isExpired()
}

func (c *Credentials) isExpired() bool {
	curTime := time.Now()
	return c.value.Expiration.Before(curTime)
}

func (c *Credentials) AccessKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value.AccessKey
}

func (c *Credentials) SecretKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value.SecretKey
}

func (c *Credentials) Valid() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value.AccessKey != "" && c.value.SecretKey != ""
}

// HasProvider reports whether a provider is set, even if retrieving failed.
func (c *Credentials) HasProvider() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.provider != nil
}

func (c *Credentials) ProviderName() string {
	if c.provider == nil {
		return "NoCredentialProvider"
	}
	return c.provider.Name()
}

// Retrieve returns a copy so the access key and secret key always match.
func (c *Credentials) Retrieve() *Credentials {
	waitStart := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.provider != nil && c.isExpired() {
		// Reuse a refresh that finished while waiting. Check the error, since
		// keys without an Expiration always look expired.
		if !c.lastRefresh.After(waitStart) {
			c.refresh()
		}
		if c.lastErr != nil {
			return nil
		}
	}
	return c.copy()
}

func (c *Credentials) refresh() {
	c.value, c.lastErr = c.provider.Retrieve()
	c.lastRefresh = time.Now()
}

func (c *Credentials) copy() *Credentials {
	return &Credentials{value: c.value, provider: c.provider}
}

// configuredProvider reports whether a provider is set up.
type configuredProvider interface {
	configured() bool
}

// LoadCredentials keeps a configured provider that fails, so it does not fall
// back to another identity.
func LoadCredentials(providers []Provider) *Credentials {
	for _, p := range providers {
		credValue, err := p.Retrieve()
		if err == nil {
			return &Credentials{value: credValue, provider: p}
		}
		if c, ok := p.(configuredProvider); ok && c.configured() {
			return &Credentials{provider: p}
		}
	}

	return &Credentials{}
}

func DefaultCredentialsChain() []Provider {
	return []Provider{
		&EnvProvider{},
		&ConfigFileProvider{},
		&PodIdentityProvider{},
		&ServerRoleProvider{
			ApiClient: metadata.NewApiClient(),
		},
	}
}

type Provider interface {
	// Retrieve returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable, or empty.
	Retrieve() (Value, error)
	// Name provider name
	Name() string
}
