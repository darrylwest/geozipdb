// simple verion
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-07-01 09:59:37

package geozipdb

import "fmt"

const (
	major = 18
	minor = 5
	patch = 26
)

// Version - return the version number as a single string
func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
