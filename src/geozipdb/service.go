//
// handlers
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-07-01 12:57:59
//

package geozipdb

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

// Keytype - keys are strings
type Keytype string

// Ziptype - zipcodes are strings
type Ziptype string

// Coord - lat/lng
type Coord struct {
	Lat float64
	Lng float64
}

// ZipcodeCoord - zipcode and coord lat/lng
type ZipcodeCoord struct {
	Zipcode Ziptype
	Coord
}

// Service - the primary service struct
type Service struct {
	config *Config
}

var keyMap = make(map[Keytype][]*ZipcodeCoord)
var zipMap = make(map[Ziptype]*Coord)
var initialized = false

// CreateKey - create a key from lat/lng
func (svc Service) CreateKey(lat, lng float64) Keytype {
	llat := int(lat * 10)
	llng := int(lng * 10)

	return Keytype(fmt.Sprintf("%d:%d", llat, llng))
}

// Initialize - initialize the data
func (svc Service) Initialize() {
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

// CoordFromZip - return the coordinate of this zipcode
func (svc Service) CoordFromZip(code Ziptype) (*Coord, bool) {
	v, ok := zipMap[code]
	return v, ok
}

// ZipListFromCoord - return a list of zip codes that are near the coordinates
func (svc Service) ZipListFromCoord(coord *Coord) ([]*ZipcodeCoord, bool) {
	key := svc.CreateKey(coord.Lat, coord.Lng)
	v, ok := keyMap[key]

	return v, ok
}

// NewService - create the service based on config
func NewService(config *Config) *Service {
	svc := new(Service)

	svc.config = config

	return svc
}

// return the zip list for a given coordinate lat/lng
func (svc Service) ziplistHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := ps.ByName("coords")
	fmt.Printf("find zip list for coords %s\n", p)

	fmt.Fprintf(w, "p %s\n\r", ps.ByName("coord"))
}

// return the coordinates for a given zip
func (svc Service) coordHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	zipcode := ps.ByName("zip")
	fmt.Printf("find coords for zip %s\n", zipcode)

	if coord, ok := svc.CoordFromZip(Ziptype(zipcode)); ok {
		fmt.Fprintf(w, "coord:%f/%f\n\r", coord.Lat, coord.Lng)
	} else {
		fmt.Fprintf(w, "not found for zip %s\n\r", zipcode)
	}
}

// Start - initialize the data and start the listener service
func (svc Service) Start() {
	if initialized == false {
		svc.Initialize()
	}

	router := httprouter.New()

	router.GET("/v1/zipdb/coord/:zip", svc.coordHandler)
	router.GET("/v1/zipdb/ziplist/:coord", svc.ziplistHandler)

	port := svc.config.Port
	host := fmt.Sprintf(":%d", port)
	fmt.Printf("listening on port %d\n", port)

	err := http.ListenAndServe(host, router)
	if err != nil {
		panic(err)
	}
}
