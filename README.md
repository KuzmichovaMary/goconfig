# Simplest Config module ever

## Installation

`go get github.com/KuzmichovaMary/goconfig`

## Usage

Suppose you have this config file:

```yaml
# config.yaml
host: localhost
port: 8080
```

```golang
import (
    "github.com/KuzmichovaMary/goconfig"
)

type ServerConfig struct {
    Host         string `yaml:host`
    Port         string `yaml:port`
}

func main() {
    var cfg ServerConfig
    goconfig.ConfigFromFile("config.yaml", &cfg)
    fmt.Print(cfg.Host) // >> localhost
    fmt.Print(cfg.Host) // >> 8080
}

```