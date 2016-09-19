FROM golang:1.7.1-alpine

COPY . $GOPATH/src/github.com/IrfanFaizullabhoy/pacific

WORKDIR $GOPATH/src/github.com/IrfanFaizullabhoy/pacific/server

RUN go build server.go

ENTRYPOINT ["go","run","server.go"]