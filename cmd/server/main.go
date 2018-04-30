package main

import (
	"context"
	"log"
	"net"

	"github.com/lpan/grpc-lame-duck/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type adderServer struct{}

func (adderServer) Add(
	ctx context.Context, in *pb.AddRequest,
) (*pb.AddReply, error) {
	var result int32
	for _, o := range in.Operands {
		result += o
	}
	return &pb.AddReply{result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("failed to listen to port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAdderServer(s, &adderServer{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
