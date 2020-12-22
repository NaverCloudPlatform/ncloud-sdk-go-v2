package credentials

import (
	"errors"
	"os"
)

type EnvProvider struct {
	retrieved bool
}

func (p *EnvProvider) Name() string {
	return "EnvProvider"
}

// Retrieve retrieves the keys from the environment.
func (p *EnvProvider) Retrieve() (Value, error) {
	p.retrieved = false
	id := os.Getenv("NCLOUD_ACCESS_KEY_ID")
	if id == "" {
		id = os.Getenv("NCLOUD_ACCESS_KEY")
	}

	secret := os.Getenv("NCLOUD_SECRET_ACCESS_KEY")
	if secret == "" {
		secret = os.Getenv("NCLOUD_SECRET_KEY")
	}

	if id == "" {
		return Value{}, errors.New("NCLOUD_ACCESS_KEY_ID or NCLOUD_ACCESS_KEY not found in environment")
	}

	if secret == "" {
		return Value{}, errors.New("NCLOUD_SECRET_ACCESS_KEY or NCLOUD_SECRET_KEY not found in environment")
	}

	p.retrieved = true
	return Value{
		AccessKey: id,
		SecretKey: secret,
	}, nil
}

// IsExpired returns if the credentials have been retrieved.
func (p *EnvProvider) IsExpired() bool {
	return !p.retrieved
}
