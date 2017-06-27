package unit

import (
	"os"
	"geozipdb"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		g.It("should create a config struct", func() {
			cfg := new(spotcache.Config)
			g.Assert(cfg.Baseport).Equal(0)
			g.Assert(cfg.Home).Equal("")
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := spotcache.NewDefaultConfig()

			g.Assert(cfg.Home).Equal(home)
			g.Assert(cfg.Port).Equal(19501)
		})
	})
}
