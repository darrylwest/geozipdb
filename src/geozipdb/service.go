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
	log.Info("initialize the database...")
	lines := strings.Split(geodata, "\n")

	log.Info("processing %d data rows...\n", len(lines))
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

	log.Info("processed %d rows...\n", len(lines))
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

func (svc Service) errorHandler(w http.ResponseWriter, r *http.Request, msg string) {
    log.Warn("%s", msg)
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "%s\n\r", msg)
}

// return the zip list for a given coordinate lat/lng
func (svc Service) ziplistHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := ps.ByName("coord")
	log.Info("find zip list for coord %s\n", p)

    fields := strings.Split(p, ",")
    if len(fields) != 2 {
        svc.errorHandler(w, r, fmt.Sprintf("coordinates lat/lng not formatted correctly from %s", p))
        return
    }

    lat, err := strconv.ParseFloat(fields[0], 64)
    if err != nil {
        svc.errorHandler(w, r, fmt.Sprintf("could not parse lat from %s", p))
        return
    }

    lng, err := strconv.ParseFloat(fields[1], 64)
    if err != nil {
        svc.errorHandler(w, r, fmt.Sprintf("could not parse lng from %s", p))
        return
    }

    log.Info("find list from coords %f,%f", lat, lng)

    coord := Coord{lat, lng}

    if list, ok := svc.ZipListFromCoord(&coord); ok == true {
        fmt.Fprintf(w, "%v\n\r", list)
    } else {
        svc.errorHandler(w, r, fmt.Sprintf("could not find zipcodes for %s", p))
        return
    }

}

// return the coordinates for a given zip
func (svc Service) coordHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	zipcode := ps.ByName("zip")
	log.Debug("find coord for zip %s\n", zipcode)

	if coord, ok := svc.CoordFromZip(Ziptype(zipcode)); ok {
		str := fmt.Sprintf("%f,%f", coord.Lat, coord.Lng)
		log.Info("found %s for zip %s", str, zipcode)
		fmt.Fprintf(w, "%s\n\r", str)
	} else {
        svc.errorHandler(w, r, fmt.Sprintf("could not find coordinates for zip %s", zipcode))
        return
	}
}

// Start - initialize the data and start the listener service
func (svc Service) Start() {
	if initialized == false {
		svc.Initialize()
	}

	cfg := svc.config

	router := httprouter.New()

	rname := fmt.Sprintf("%s/coord/:zip", cfg.PrimaryRoute)
	router.GET(rname, svc.coordHandler)
	log.Info("added route %s\n", rname)

	rname = fmt.Sprintf("%s/ziplist/:coord", cfg.PrimaryRoute)
	router.GET(rname, svc.ziplistHandler)
	log.Info("added route %s\n", rname)

	port := svc.config.Port
	host := fmt.Sprintf(":%d", port)
	log.Info("listening on port %d\n", port)

	err := http.ListenAndServe(host, router)
	if err != nil {
		panic(err)
	}
}
