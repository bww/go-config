package staticsecrets

import (
	"github.com/bww/go-config/v1/secrets"
)

type Postproc struct {
	secrets map[string]string
}

func New(secrets map[string]string) *Postproc {
	return &Postproc{secrets}
}

func (p *Postproc) Unwrap(name, value string) (string, error) {
	if sec, ok := p.secrets[value]; ok {
		return sec, nil
	} else {
		return "", secrets.ErrNotFound
	}
}
