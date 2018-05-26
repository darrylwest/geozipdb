//
// status
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-07-03 11:37:13
//

package geozipdb

import (
	"encoding/json"
	"runtime"
    "os"
	"time"
)

var started = time.Now().Unix()

// Status - the standard status struct
type Status struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
    PID       int    `json:"pid"`
	CPUs      int    `json:"cpus"`
	GoVers    string `json:"go"`
	TimeStamp int64  `json:"ts"`
	UpTime    int64  `json:"uptime-seconds"`
}

// GetStatus return the current status struct
func GetStatus() Status {
	now := time.Now().Unix()

	s := Status{}
	s.Status = "ok"
	s.Version = Version()
    s.PID = os.Getpid()
	s.CPUs = runtime.NumCPU()
	s.GoVers = runtime.Version()
	s.TimeStamp = now
	s.UpTime = now - started

	return s
}

// GetStatusAsJSON return the current status as a json string
func GetStatusAsJSON() string {
	status := GetStatus()
	json, _ := json.Marshal(status)

	return string(json)
}
