package staticsecrets

import (
	"github.com/bww/go-config/v1/secrets"
)

type Provider struct {
	secrets map[string]string
}

func New(secrets map[string]string) *Provider {
	return &Provider{secrets}
}

func (p *Provider) Unwrap(name, value string) (string, error) {
	if sec, ok := p.secrets[value]; ok {
		return sec, nil
	} else {
		return "", secrets.ErrNotFound
	}
}
