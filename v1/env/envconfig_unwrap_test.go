package env

import (
	"fmt"
	"os"
	"testing"

	"github.com/bww/go-config/v1/secrets/staticsecrets"
	"github.com/stretchr/testify/assert"
)

var secretPostproc = staticsecrets.New(map[string]string{
	"value_1": "unwrapped_value_1",
	"value_2": "unwrapped_value_2",
})

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
	p.RegisterSecrets(secretPostproc)
	err := p.Process("env", &s)
	assert.Nil(t, err, fmt.Sprint(err))

	assert.Equal(t, "unwrapped_value_1", s.Value1)
	assert.Equal(t, "unwrapped_value_2", s.Value2)
	assert.Equal(t, "value_3", s.Value3)
}
