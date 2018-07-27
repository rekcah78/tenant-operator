FROM alpine:3.6

RUN adduser -D tenant-operator
USER tenant-operator

ADD tmp/_output/bin/tenant-operator /usr/local/bin/tenant-operator
