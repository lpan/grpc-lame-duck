package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lpan/grpc-lame-duck/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("dns:///localhost:1337", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	adder := pb.NewAdderClient(conn)
	operands := []int32{1, 2, 3}

	result, err := adder.Add(context.TODO(), &pb.AddRequest{operands})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
