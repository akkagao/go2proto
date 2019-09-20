package main

// server.go

import (
	"encoding/json"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"go2protoexample/proto"
	"go2protoexample/util"
)

const (
	port = ":50051"
)

type DemoService struct{}

func (s *DemoService) CreateDemo(ctx context.Context, demoRequest *proto.DemoRequest) (*proto.DemoResponse, error) {
	j, _ := json.Marshal(demoRequest)
	log.Println(string(j))
	return util.CreateResult(demoRequest), nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterDemoServer(s, &DemoService{})
	s.Serve(lis)
}
