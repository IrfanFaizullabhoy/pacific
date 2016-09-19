package main

import (
	"fmt"

	pb "github.com/IrfanFaizullabhoy/pacific/schema"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// dials connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	// closes the connection at the end of main()
	defer conn.Close()

	// creates client, this is used as argument in requests
	client := pb.NewPassObjectClient(conn)

	// explicit definition of object I am sending to server
	object := pb.Object{
		Name: "firstObject",
		Size: 1,
	}

	//send object to the server
	_, err = client.SendObject(context.Background(), &object)

	check(err)

	// use object name to receive object
	_, err1 := client.ReceiveObject(context.Background(), &pb.ObjectRequest{
		Name: object.Name,
	})
	check(err1)

	fmt.Println("Success, received the following object")
}
