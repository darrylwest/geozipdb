package unit

import (
	"fmt"
	"geozipdb"
	"testing"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)
	cfg := new(geozipdb.Config)

	g.Describe("Service", func() {
		g.It("should create a service struct", func() {
			service := geozipdb.NewService(cfg)
			g.Assert(fmt.Sprintf("%T", service)).Equal("*geozipdb.Service")
		})

		g.It("should initialize and start the service", func() {
			service := geozipdb.NewService(cfg)
			service.Initialize()

			zip := geozipdb.Ziptype("94705")
			coord, ok := service.CoordFromZip(zip)
			g.Assert(ok).IsTrue()
			g.Assert(coord.Lat > 0).IsTrue()
			g.Assert(coord.Lng < 0).IsTrue()
		})

		g.It("should find a list of zip codes for a specific coordinate", func() {
			service := geozipdb.NewService(cfg)
			service.Initialize()

			zip := geozipdb.Ziptype("94705")
			coord, _ := service.CoordFromZip(zip)

			list, ok := service.ZipListFromCoord(coord)
			g.Assert(ok).IsTrue()

			g.Assert(len(list)).Equal(30)
		})
	})
}
