package credentials

import (
	"crypto/rand"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud/metadata"
	"math/big"
	"sync"
	"time"
)

const (
	// Refresh window before expiry for keys with an Expiration. Exchange
	// servers may hand the same cached keys to many clients, so a random
	// point spreads their refreshes.
	refreshWindowMin = 7 * time.Minute
	refreshWindowMax = 15 * time.Minute

	// Minimum interval between refreshes while the cached keys are valid.
	refreshRetryInterval = time.Minute
)

type Credentials struct {
	mu       sync.RWMutex
	value    Value
	provider Provider
	// refreshAt is zero when refreshing only on expiry.
	refreshAt   time.Time
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

	if c.provider != nil && c.needsRefresh(time.Now()) {
		// Reuse a refresh that finished while waiting. Check the error, since
		// keys without an Expiration always look expired.
		if !c.lastRefresh.After(waitStart) {
			c.refresh()
		}
		if c.lastErr != nil && c.isExpired() {
			return nil
		}
	}
	return c.copy()
}

func (c *Credentials) needsRefresh(now time.Time) bool {
	return c.isExpired() || (!c.refreshAt.IsZero() && !now.Before(c.refreshAt))
}

func (c *Credentials) refresh() {
	value, err := c.provider.Retrieve()
	now := time.Now()
	c.lastRefresh, c.lastErr = now, err
	if err != nil {
		// Keep using the cached keys and retry later.
		c.refreshAt = now.Add(refreshRetryInterval)
		return
	}
	c.setValue(value, now)
}

func (c *Credentials) copy() *Credentials {
	return &Credentials{value: c.value, provider: c.provider}
}

func (c *Credentials) setValue(value Value, now time.Time) {
	c.value = value
	c.refreshAt = time.Time{}

	if value.Expiration.IsZero() {
		return
	}
	// The server may return cached keys with less time left than the window.
	floor := now.Add(refreshRetryInterval)
	c.refreshAt = value.Expiration.Add(-randomDuration(refreshWindowMin, refreshWindowMax))
	if c.refreshAt.Before(floor) {
		c.refreshAt = floor
	}
}

// The math/rand global source is not seeded before Go 1.20.
func randomDuration(min, max time.Duration) time.Duration {
	if max <= min {
		return min
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)+1))
	if err != nil {
		return max
	}
	return min + time.Duration(n.Int64())
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
			credentials := &Credentials{provider: p}
			credentials.setValue(credValue, time.Now())
			return credentials
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
