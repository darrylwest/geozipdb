package unit

import (
	"geozipdb"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		g.It("should create a config struct", func() {
			cfg := new(geozipdb.Config)
			g.Assert(cfg.Port).Equal(0)
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := geozipdb.NewDefaultConfig()

			g.Assert(cfg.Port).Equal(29444)
		})
	})
}
