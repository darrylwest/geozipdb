#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2017-07-17 08:28:19
#

network=service-net
docker network create $network

image=ebay-local/geozipdb:latest

docker run --name geozipdb-1 -d -p 4540:4539 --network=$network $image 
docker run --name geozipdb-2 -d -p 4541:4539 --network=$network $image 
docker run --name geozipdb-3 -d -p 4542:4539 --network=$network $image 

docker ps | fgrep geozipdb-
