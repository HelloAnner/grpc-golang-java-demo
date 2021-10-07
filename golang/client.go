package main

import (
	"context"
	"fmt"
	"grpc-demo/hello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connect server fail : %s\n", err)
		return
	}

	defer conn.Close()

	// create client
	c := hello.NewHelloClient(conn)

	response, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "golang client"})
	if err != nil {
		fmt.Printf("call hello service fail: %s\n", err)
		return
	}

	fmt.Printf("response : %s\n", response.Message)
}
