#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2017-07-17 08:28:19
#

ports="4540 4541 4542"
image=ebay-local/geozipdb:latest

for p in $ports
do
    # echo $p
    docker run --name geozipdb-$p -d -p $p:4539 --network=ebay-local $image 
done


docker ps | fgrep geozipdb-
