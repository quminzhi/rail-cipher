FROM golang:bookworm

LABEL rail-version="1.0.0"
MAINTAINER mqu@nyit.edu

WORKDIR /workspace

COPY ./rail.go ./rail.go
COPY ./go.mod ./go.mod
COPY ./rail_test.go ./rail_test.go

# Add a running foreground process
CMD ["/usr/bin/tail", "-f", "/etc/default/useradd"]