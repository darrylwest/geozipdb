// logger
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 10:57:53

package geozipdb

import (
	"fmt"
	"os"

	"github.com/darrylwest/cassava-logger/logger"
)

var log *logger.Logger

// CreateLogger create a new console logger; use log.SetLevel( logger.WarnLevel )
func CreateLogger() *logger.Logger {
	if log == nil {
		handler, err := logger.NewStreamHandler(os.Stdout)

		if err != nil {
			fmt.Printf("%s\n", err)
			panic("logger could not be created...")
		}

		log = logger.NewLogger(handler)
	}

	return log
}

