//
// handlers
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-07-01 12:57:59
//

package geozipdb

import (
	"fmt"
	"strconv"
	"strings"
)

type Keytype string
type Ziptype string

type Coord struct {
	Lat float64
	Lng float64
}

type ZipcodeCoord struct {
	Zipcode Ziptype
	Coord
}

type Service struct {
}

var keyMap = make(map[Keytype][]*ZipcodeCoord)
var zipMap = make(map[Ziptype]*Coord)

func (svc Service) CreateKey(lat, lng float64) Keytype {
	llat := int(lat * 10)
	llng := int(lng * 10)

	return Keytype(fmt.Sprintf("%d:%d", llat, llng))
}

func (svc Service) initialize() {
	fmt.Println("initialize the database...")
	lines := strings.Split(geodata, "\n")

	fmt.Printf("processing %d data rows...\n", len(lines))
	for i := 0; i < len(lines); i++ {
		fields := strings.Split(lines[i], ",")
		zipcode := Ziptype(fields[0])

		lat, _ := strconv.ParseFloat(fields[1], 64)
		lng, _ := strconv.ParseFloat(fields[2], 64)

		key := svc.CreateKey(lat, lng)

		coord := Coord{lat, lng}
		zipMap[zipcode] = &coord

		zcoord := ZipcodeCoord{zipcode, coord}
		keyMap[key] = append(keyMap[key], &zcoord)
	}

	fmt.Printf("processed %d rows...\n", len(lines))
}

func (svc Service) CoordFromZip(code Ziptype) (*Coord, bool) {
	v, ok := zipMap[code]
	return v, ok
}

func (svc Service) ZipListFromCoord(coord *Coord) ([]*ZipcodeCoord, bool) {
	key := svc.CreateKey(coord.Lat, coord.Lng)
	v, ok := keyMap[key]

	return v, ok
}

func (svc Service) Start() {
	svc.initialize()
}

func NewService(config *Config) *Service {
	svc := new(Service)

	return svc
}
