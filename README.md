# Configuration
This package loads configuration data from the environment. It is based on the excellent [`envconfig`](https://github.com/kelseyhightower/envconfig) package, which you should probably use instead unless you specifically need the post-processing functionality offered by `go-config`.

## Unwrapping Secrets
This package extends `envconfig` to include support for post-processing configuration data loaded from the environment before it is assigned to a field. The specific intention of this functionality is to support unwrapping secrets that are provided externally.

Let's say the environment contains:

```sh
EXAMPLE_URL=https://github.com/bww/go-config
EXAMPLE_PASSWORD=password_name
```

We might load this confirmation using the following illustrative snippet:

```go
import (
  "github.com/bww/go-config/v1/env"
  "github.com/bww/go-config/v1/secrets/staticsecrets"
)

type Config struct {
  URL      string `env:"URL"`                      // process this field normally
  Password string `env:"PASSWORD" unwrap:"secret"` // unwrap this field as a secret
}

var secrets = staticsecrets.New(map[string]string{
  "password_name": "password_secret",
})

proc := &env.Processor{}
proc.RegisterSecrets(secrets)

conf := Config{}
err := proc.Process("example", &conf)
if err != nil {
  panic(err)
}

fmt.Println(conf.URL)      // "https://github.com/bww/go-config"
fmt.Println(conf.Password) // "password_secret"
```
