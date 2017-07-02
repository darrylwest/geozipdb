//
// config  - application specification and CLI parsing
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-06-26 17:56:46

package geozipdb

import (
	"flag"
	"os"
	"path"
)

// Config the config structure
type Config struct {
	Port         int
	PrimaryRoute string
}

// NewDefaultConfig default settings
func NewDefaultConfig() *Config {
	cfg := new(Config)

	cfg.Port = 5000
	cfg.PrimaryRoute = "/v1/zipdb"

	return cfg
}

// ParseArgs parse the command line args
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")

	port := flag.Int("port", dflt.Port, "set the server's port number (e.g., 29444)...")
	route := flag.String("route", dflt.PrimaryRoute, "set the server's primary route (e.g., /v1/zipdb)...")
    level := flag.Int("loglevel", 2, "set the server's log level 0..5 for trace..error, default info=2")

	flag.Parse()

	log.Info("%s Version: %s\n", path.Base(os.Args[0]), Version())

	if *vers == true {
		os.Exit(0)
	}

	cfg := new(Config)

	cfg.Port = *port
	cfg.PrimaryRoute = *route

    log.SetLevel(*level)

	return cfg
}

// IsProduction return true if the current env is production
func IsProduction(env string) bool {
	return true
}
