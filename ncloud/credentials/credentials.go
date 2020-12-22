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

func LoadCredentials(providers []Provider) *Credentials {
	var credentials Credentials
	for _, p := range providers {
		credValue, err := p.Retrieve()
		if err == nil {
			credentials = Credentials{
				value: credValue,
				provider: p,
			}
			break
		}
	}

	return &credentials
}

func DefaultCredentialsChain() []Provider {
	return []Provider{
		&EnvProvider{},
		&ConfigFileProvider{},
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
