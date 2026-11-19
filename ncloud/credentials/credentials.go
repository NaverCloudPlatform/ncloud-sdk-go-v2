package credentials

import (
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud/metadata"
	"time"
)

type Credentials struct {
	value    Value
	provider Provider
}

type Value struct {
	AccessKey  string
	SecretKey  string
	Expiration time.Time
}

// IsExpired returns if the credentials are expired.
func (c *Credentials) IsExpired() bool {
	curTime := time.Now()
	return c.value.Expiration.Before(curTime)
}

func (c *Credentials) AccessKey() string {
	return c.value.AccessKey
}

func (c *Credentials) SecretKey() string {
	return c.value.SecretKey
}

func (c *Credentials) Valid() bool {
	return c.value.AccessKey != "" && c.value.SecretKey != ""
}

// HasProvider reports whether a provider is set, even if retrieving failed.
func (c *Credentials) HasProvider() bool {
	return c.provider != nil
}

func (c *Credentials) ProviderName() string {
	if c.provider == nil {
		return "NoCredentialProvider"
	}
	return c.provider.Name()
}

func (c *Credentials) Retrieve() *Credentials {
	if c.provider != nil && c.IsExpired() {
		var err error
		if c.value, err = c.provider.Retrieve(); err != nil {
			return nil
		}
	}
	return c
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
