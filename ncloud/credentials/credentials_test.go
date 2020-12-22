package credentials

import (
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
