//
// config  - application specification and CLI parsing
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-06-26 17:56:46

package geozipdb

import (
	"flag"
	"fmt"
	"os"
	"path"
)

// Config the config structure
type Config struct {
	Port int
}

// NewDefaultConfig default settings
func NewDefaultConfig() *Config {
	cfg := new(Config)

	cfg.Port = 29444

	return cfg
}

// ParseArgs parse the command line args
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")

	port := flag.Int("port", dflt.Port, "set the server's port number (e.g., 29444)...")

	flag.Parse()

	fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), Version())

	if *vers == true {
		os.Exit(0)
	}

	cfg := new(Config)

	cfg.Port = *port

	return cfg
}

// IsProduction return true if the current env is production
func IsProduction(env string) bool {
	return true
}
