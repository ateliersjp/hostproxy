FROM golang:alpine AS build

ARG IGNORECACHE=0

ADD . /hostproxy
RUN target=/go \
    cd /hostproxy \
    && echo "go get" \
    && go get \
    && echo "go build" \
    && GOCACHE=/go/.cache CGO_ENABLED=0 go build -ldflags='-s -w'

FROM alpine:latest

COPY --from=build /hostproxy/hostproxy /bin/
COPY ./net.ipv6.conf /etc/sysctl.d/
COPY ./start.sh /bin/

ENTRYPOINT [ "start.sh" ]
