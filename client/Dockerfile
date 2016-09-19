FROM golang:1.7.1-alpine

COPY . $GOPATH/src/github.com/IrfanFaizullabhoy/pacific

WORKDIR $GOPATH/src/github.com/IrfanFaizullabhoy/pacific/client

RUN go build client.go

ENTRYPOINT ["go","run","client.go"]