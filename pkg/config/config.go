package config

import (
	"fmt"
	"net/url"

	"github.com/l-vitaly/goenv"
)

var urlNil = url.URL{}

const errPattern = "could not set %s"

// env name constants
const (
	DbConnEnvName = "MIGRATION_DB_CONN"
)

// Config config.
type Config struct {
	DbConn url.URL
}

// Parse env config vars.
func Parse() (*Config, error) {
	cfg := &Config{}

	goenv.URLVar(&cfg.DbConn, DbConnEnvName, url.URL{})

	goenv.Parse()

	if cfg.DbConn == urlNil {
		return nil, fmt.Errorf(errPattern, DbConnEnvName)
	}

	return cfg, nil
}
