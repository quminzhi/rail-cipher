FROM golang:bookworm

LABEL rail-version="1.0.0"
MAINTAINER mqu@nyit.edu

WORKDIR /workspace

COPY ./main ./main
COPY ./rail ./rail
RUN cd ./main && go build -o ../cipher .

# Add a running foreground process
CMD ["/usr/bin/tail", "-f", "/etc/default/useradd"]