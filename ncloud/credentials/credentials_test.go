package credentials

import (
	"errors"
	"sync"
	"sync/atomic"
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

// blockingProvider holds the first call until release is closed.
type blockingProvider struct {
	value   Value
	err     error
	calls   int32
	started chan struct{}
	release chan struct{}
}

func newBlockingProvider(value Value, err error) *blockingProvider {
	return &blockingProvider{value: value, err: err, started: make(chan struct{}), release: make(chan struct{})}
}

func (p *blockingProvider) Name() string {
	return "blockingProvider"
}

func (p *blockingProvider) Retrieve() (Value, error) {
	if atomic.AddInt32(&p.calls, 1) == 1 {
		close(p.started)
	}
	<-p.release
	return p.value, p.err
}

// retrieveConcurrently starts one Retrieve, then n more while it is in flight.
func retrieveConcurrently(creds *Credentials, p *blockingProvider, n int) []*Credentials {
	results := make([]*Credentials, n+1)
	var wg sync.WaitGroup
	retrieve := func(i int) {
		defer wg.Done()
		results[i] = creds.Retrieve()
	}
	wg.Add(1)
	go retrieve(0)
	<-p.started
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go retrieve(i)
	}
	// Let the other calls block on the refresh in flight.
	time.Sleep(50 * time.Millisecond)
	close(p.release)
	wg.Wait()
	return results
}

func TestRetrieveRefreshesOnceForConcurrentCalls(t *testing.T) {
	p := newBlockingProvider(Value{AccessKey: "new", SecretKey: "new", Expiration: time.Now().Add(time.Hour)}, nil)
	creds := &Credentials{
		value:    Value{AccessKey: "old", SecretKey: "old", Expiration: time.Now().Add(-time.Minute)},
		provider: p,
	}

	for _, got := range retrieveConcurrently(creds, p, 9) {
		if got == nil || got.AccessKey() != "new" {
			t.Fatalf("Expected: new, Actual: %v", got)
		}
	}
	if n := atomic.LoadInt32(&p.calls); n != 1 {
		t.Fatalf("Expected: 1 provider call, Actual: %d", n)
	}
}

func TestRetrieveSharesFailedRefresh(t *testing.T) {
	p := newBlockingProvider(Value{}, errors.New("provider unavailable"))
	creds := &Credentials{
		value:    Value{AccessKey: "old", SecretKey: "old", Expiration: time.Now().Add(-time.Minute)},
		provider: p,
	}

	for _, got := range retrieveConcurrently(creds, p, 9) {
		if got != nil {
			t.Fatalf("Expected: nil, Actual: %v", got)
		}
	}
	if n := atomic.LoadInt32(&p.calls); n != 1 {
		t.Fatalf("Expected: 1 provider call, Actual: %d", n)
	}

	// A later call tries the provider again.
	creds.Retrieve()
	if n := atomic.LoadInt32(&p.calls); n != 2 {
		t.Fatalf("Expected: 2 provider calls, Actual: %d", n)
	}
}

func TestRetrieveReturnsCopy(t *testing.T) {
	p := &fakeProvider{value: Value{AccessKey: "new", SecretKey: "new", Expiration: time.Now().Add(time.Hour)}}
	creds := &Credentials{
		value:    Value{AccessKey: "old", SecretKey: "old", Expiration: time.Now().Add(time.Hour)},
		provider: p,
	}

	got := creds.Retrieve()
	creds.value = Value{AccessKey: "rotated", SecretKey: "rotated", Expiration: time.Now().Add(time.Hour)}
	if got.AccessKey() != "old" || got.SecretKey() != "old" {
		t.Fatalf("Expected: old/old, Actual: %s/%s", got.AccessKey(), got.SecretKey())
	}
}

func TestRetrieveConcurrentCallsWithoutExpiration(t *testing.T) {
	// EnvProvider, ConfigFileProvider and ValueProvider return no Expiration,
	// so their keys always look expired.
	p := newBlockingProvider(Value{AccessKey: "AK", SecretKey: "SK"}, nil)
	creds := &Credentials{value: Value{AccessKey: "AK", SecretKey: "SK"}, provider: p}

	for _, got := range retrieveConcurrently(creds, p, 9) {
		if got == nil || got.AccessKey() != "AK" {
			t.Fatalf("Expected: AK, Actual: %v", got)
		}
	}
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
