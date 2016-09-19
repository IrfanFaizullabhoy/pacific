FROM golang:1.7

COPY . $GOPATH/src/github.com/IrfanFaizullabhoy/pacific

WORKDIR $GOPATH/src/github.com/IrfanFaizullabhoy/pacific/client

##RUN go get google.golang.org/grpc
##RUN go get golang.org/x/net/context
##RUN go get github.com/golang/protobuf/proto

RUN go build client.go

ENTRYPOINT ["go","run","client.go"]