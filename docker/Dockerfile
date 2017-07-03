#
# @author darryl.west <darryl.west@raincitysoftware.com>
# @created 2017-07-03 10:15:31
#

FROM scratch

LABEL maintainer="darryl.west@raincitysoftware.com"

COPY ./geozipdb /geozipdb

EXPOSE 4539

# uses the default port 5000 and route /v1/zipdb
CMD ["/geozipdb"]
