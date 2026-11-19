package credentials

import (
	"errors"
	"testing"
	"time"
)

func TestIsExpired(t *testing.T) {
	creds := &Credentials{
		value: Value{
			Expiration: time.Now().Add(time.Minute),
		},
	}
	if creds.IsExpired() {
		t.Fatalf("Expected: false, Actual: %v", creds.IsExpired())
	}
}

type fakeProvider struct {
	value Value
	err   error
}

func (p *fakeProvider) Name() string {
	return "fakeProvider"
}

func (p *fakeProvider) Retrieve() (Value, error) {
	return p.value, p.err
}

type configuredFakeProvider struct {
	fakeProvider
}

func (p *configuredFakeProvider) configured() bool {
	return true
}

func TestLoadCredentialsKeepsConfiguredProvider(t *testing.T) {
	first := &configuredFakeProvider{fakeProvider{err: errors.New("exchange failed")}}
	second := &fakeProvider{value: Value{AccessKey: "fallback", SecretKey: "fallback"}}

	creds := LoadCredentials([]Provider{first, second})

	if creds.provider != first || creds.Valid() {
		t.Fatalf("Expected: empty credentials from the first provider, Actual: %s, valid %v", creds.ProviderName(), creds.Valid())
	}

	first.value, first.err = Value{AccessKey: "AK", SecretKey: "SK", Expiration: time.Now().Add(time.Hour)}, nil
	if got := creds.Retrieve(); got == nil || got.AccessKey() != "AK" {
		t.Fatalf("Expected: AK, Actual: %v", got)
	}
}

func TestLoadCredentialsSkipsUnconfiguredProvider(t *testing.T) {
	first := &fakeProvider{err: errors.New("not set")}
	second := &fakeProvider{value: Value{AccessKey: "AK", SecretKey: "SK"}}

	creds := LoadCredentials([]Provider{first, second})

	if creds.provider != second || creds.AccessKey() != "AK" {
		t.Fatalf("Expected: AK from the second provider, Actual: %s", creds.AccessKey())
	}
}
