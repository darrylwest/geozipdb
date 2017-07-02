package main

import (
	"fmt"
	"geozipdb"
)

func main() {
	geozipdb.CreateLogger()
	config := geozipdb.ParseArgs()

	service := geozipdb.NewService(config)

	fmt.Printf("version %s\n", geozipdb.Version())
	service.Start()
}
