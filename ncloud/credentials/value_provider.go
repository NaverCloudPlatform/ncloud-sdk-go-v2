package credentials

type ValueProvider struct {
	value Value
}

func NewValueProviderCreds(accessKey, secretKey string) *Credentials {
	p := NewValueProvider(accessKey, secretKey)
	return &Credentials{
		value:    p.value,
		provider: p,
	}
}

func NewValueProvider(accessKey, secretKey string) *ValueProvider {
	return &ValueProvider{
		value: Value{
			AccessKey:  accessKey,
			SecretKey:  secretKey,
		},
	}
}

func (p *ValueProvider) Name() string {
	return "ValueProvider"
}

func (p *ValueProvider) Retrieve() (Value, error) {
	return p.value, nil
}
