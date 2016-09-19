package main;

import (
	"fmt"
	pb "github.com/IrfanFaizullabhoy/pacific/schema"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":8080"
)

type passObjectServer struct{}


// runs whenever request `SendObject` is received
func (server *passObjectServer) SendObject(ctx context.Context, in *pb.Object) (*pb.Response, error) {
	// add object to database
	fmt.Println("Received the object: ", in.Name)
	response := pb.Response {
		Response: in.Name,
	}
	return &response, nil
}

// runs whenever request `ReceiveObject` is received
func (server *passObjectServer) ReceiveObject(ctx context.Context, in *pb.ObjectRequest) (*pb.Object, error) {
	// query database for object
	object := pb.Object {
		Name: in.Name,
		Size: 1,
	}
	return &object, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("EKFLEJFLEJF")
	}

	s := grpc.NewServer()
	pb.RegisterPassObjectServer(s, &passObjectServer{})
	s.Serve(lis)
}