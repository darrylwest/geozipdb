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

* Clone from [github](https://github.com/darrylwest/geozipdb).  Type `make install-deps build` and look in the bin folder
* go get github.com/darrywest/geozipdb 
* download from docker repo (coming soon...)

## Use

### As a Service

You can run geozipdb as a stand-alone binary or as a docker container.  The container size is less than 7Mb and built from _scratch_.

Assuming the service is running on the default port 5000 with the default route of /v1/zipdb to get the lat/lng coordinates of zipcode 94705:

`curl http://localhost:5000/v1/zipdb/coord/94705 -> 37.865183,-122.238209`

A single response with the lat,lng is returned.  A 404 is returned if the zipcode is not in out database.

To get a list of zip codes from coordinates do this:

`curl http://localhost:5000/v1/zipdb/ziplist/37.865183,-122.238209 -> 94602,94608,94609,94610,94611...`

A comma delimited list is return or a 404 if the lat/lng does not map to a zip.

_Note: be careful not to have spaces in your url._

### As an API

config := geozipdb.NewDefaultConfig()
service := geozipdb.NewService(config)

coord := geozipdb.CoordFromZip("94705")

fmt.Println(coord.Lat, coord.Lng)

### Logging

You may change the logging level either through the command line, or in the API or with a service call.

## License

Apache 2.0

###### Copyright © 2017, Rain City Software | darryl.west | Version 18.5.26

