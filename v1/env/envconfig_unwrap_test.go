package env

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/bww/go-config/v1/secrets"
	"github.com/bww/go-config/v1/secrets/staticsecrets"
	"github.com/stretchr/testify/assert"
)

type SpecificationWithUnwrap struct {
	Value1 string `env:"VALUE_1" unwrap:"secret"`
	Value2 string `env:"VALUE_2" unwrap:"secret"`
	Value3 string `env:"VALUE_3"`
}

func TestProcessUnwrap(t *testing.T) {
	var s SpecificationWithUnwrap

	os.Clearenv()
	os.Setenv("ENV_VALUE_1", "value_1")
	os.Setenv("ENV_VALUE_2", "value_2")
	os.Setenv("ENV_VALUE_3", "value_3")

	p := &Processor{}
	p.RegisterSecrets(staticsecrets.New(map[string]string{
		"value_1": "unwrapped_value_1",
		"value_2": "unwrapped_value_2",
	}))

	err := p.Process("env", &s)
	assert.Nil(t, err, fmt.Sprint(err))

	assert.Equal(t, "unwrapped_value_1", s.Value1)
	assert.Equal(t, "unwrapped_value_2", s.Value2)
	assert.Equal(t, "value_3", s.Value3)
}

func TestProcessUnwrapError(t *testing.T) {
	var s SpecificationWithUnwrap

	os.Clearenv()
	os.Setenv("ENV_VALUE_1", "value_1")
	os.Setenv("ENV_VALUE_2", "value_2")
	os.Setenv("ENV_VALUE_3", "value_3")

	p := &Processor{}
	p.RegisterSecrets(staticsecrets.New(map[string]string{
		"value_1": "unwrapped_value_1",
		// value_2 is not present, cannot unwrap
	}))

	err := p.Process("env", &s)
	if assert.NotNil(t, err, "Expcted an error") {
		assert.Equal(t, true, errors.Is(err, secrets.ErrNotFound), "Expected not-found error")
	}

}
