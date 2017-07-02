# Geo Zip Db

```
  ___           _____        ___  _    
 / __|___ ___  |_  (_)_ __  |   \| |__ 
| (_ / -_) _ \  / /| | '_ \ | |) | '_ \
 \___\___\___/ /___|_| .__/ |___/|_.__/
                     |_|               
```

# Overview

A simple application that returns a list of US zip codes for a given latitude/longitude and optional radius.  It will also return the lat/lng coordinates for a specified US zip code.

## Installation

Here are the options:

* Clone from [github](https://github.com/darrylwest/geozipdb).  Type make build and look in the bin folder
* go get github.com/darrywest/geozipdb 
* download from docker repo (coming soon...)

## Use

### As a Service

Assuming the service is running on the default port 5000 with the default route of /v1/zipdb...

curl http://localhost:5000/v1/zipdb/coord/94705 -> 37.865183,-122.238209
curl http://localhost:5000/v1/zipdb/ziplist/37.865183,-122.238209 -> [ list ]

_Note: be careful not to have spaces in your url_

### As an API

config := geozipdb.NewDefaultConfig()
service := geozipdb.NewService(config)

coords := geozipdb.CoordFromZip("94705")

fmt.Println(coords.Lat, coords.Lng)

## License

Apache 2.0

###### Copyright © 2017, Rain City Software | darryl.west | Version 1.0.5

