package main

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"go2protoexample/proto"
	"go2protoexample/util"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect ", err)
	}
	defer conn.Close()

	c := proto.NewDemoClient(conn)
	demoResponse, err := c.CreateDemo(context.Background(), util.CreateParameter())

	if err != nil {
		log.Fatal("could not greet ", err)
	}
	j, _ := json.Marshal(demoResponse)
	log.Printf("demoResponse: %v", string(j))
}
