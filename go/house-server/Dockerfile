# Start from a Debian image with Go v12 installed and a workspace (GOPATH) configured at /go.
FROM golang:1.12
LABEL maintainer="Robbie Heywood <r.a.heywood@btinternet.com>"

ENV GO111MODULE=on

ADD . /go/src/github.com/robbieheywood/house-app/house-server

WORKDIR /go/src/github.com/robbieheywood/house-app/house-server

RUN go mod download

RUN go install github.com/robbieheywood/house-app/house-server

ENTRYPOINT /go/bin/house-server

EXPOSE 8080
