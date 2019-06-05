package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"gitlab.com/atreya2011/grpc-practice/basic-crud/basiccrud"

	"google.golang.org/grpc"
)

var addrFlag = flag.String("addr", "localhost:6000", "server address host:post")

func main() {
	// Connect with the grpc server listening on port 5000 using an insecure connection
	// This creates a new connection
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())
	// Handle the error as usual
	if err != nil {
		log.Fatalln(err)
	}
	// Close the connection
	defer conn.Close()

	client := basiccrud.NewBasicCrudClient(conn)

	// call the Read method from the BasicCrud service
	res, err := client.Read(context.Background(), &basiccrud.ReadRequest{Id: 3})
	fmt.Printf("Response : %v\n", res.GetFullname())
}
