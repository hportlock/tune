# Tune

Tune is a simple golang library that helps with loading TOML configuration
files.

Installation:

```bash
go get github.com/hportlock/tune
```

## Loading configuration

Assuming that you have a directory called `config` that contains the following
`development.toml`:

```toml
# development.toml

DatabaseURL = "foo"
```

The you could load it with the following go code:

```go
package main

import (
  "fmt"

  "github.com/hportlock/tune"
)

type Config struct {
  DatabaseURL string
}

func main() {
  config := Config{}
  configRoot := "config"
  err := tune.LoadConfig(configRoot, "development", &config)
  if err != nil {
    panic("Error loading config")
  }

  fmt.Print(config.DatabaseURL)
}
```

## Local overrides

If you have some configuration settings that you would like to override on your
local development machine the you can specify them in a file named
`<environment>.local.toml`.

For example if you had the following files:
```toml
# development.toml

DatabaseServer = "devserver.mycompany.com"
DatabaseName = "development"
```

```toml
#development.local.toml

DatabaseServer = "localhost"
```

Then after calling LoadConfig your Config struct will contain a DatabaseServer of
"localhost" and a DatabaseName of "development".

## Shared settings between environments

To have some settings that are shared across environments simply add them
to a settings.toml file in the root of you config directory.

## Environment Variables

You can also use environment variables in your TOML templates. Just use the
following syntax:

```toml
DatabasePassword = "{{(.Getenv "DBPASS")}}"
```

This will pull the environment variable "DBPASS" into the DatabasePassword
field of your config struct.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
