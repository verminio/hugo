package provider

import "net/url"

var providers = make(map[string]Provider)

type Provider interface {
	Publish(target *url.URL) error
}

func registerProvider(name string, provider Provider) {
	if provider == nil {
		panic("Provider is empty.")
	}

	if _, dup := providers[name]; dup {
		panic("Provider already registered: " + name)
	}

	providers[name] = provider
}

func GetProvider(name string) Provider {
	provider, err := providers[name]

	if !err {
		panic("Unknown provider: " + name)
	}

	return provider
}
