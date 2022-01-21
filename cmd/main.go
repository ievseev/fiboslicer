package main

import (
	"context"
	pb "fiboslicer/pkg/github.com/ievseev/fiboslicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type FiboSlicerServer struct {
	pb.UnimplementedFiboSlicerServer
}

func (s *FiboSlicerServer) GetFibonacciSlice(ctx context.Context, req *pb.SliceRequest) (*pb.FibonacciSliceResponse, error) {
	sequenceLen := req.Y + 1
	fibonacciSequence := getFibonacciSequence(sequenceLen)
	result := getFibonacciSlice(fibonacciSequence, req.X, req.Y)
	return &pb.FibonacciSliceResponse{Members: result}, nil
}

func getFibonacciSequence(n int32) *[]int32 {

	//TODO handle error: n<1

	if n == 1 {
		return &[]int32{1}
	}

	fibonacciSeq := []int32{1, 1}

	if n == 2 {
		return &fibonacciSeq
	}

	for i := 2; int32(i) < n; i++ {
		fibonacciSeq = append(fibonacciSeq, fibonacciSeq[i-1]+fibonacciSeq[i-2])
	}

	return &fibonacciSeq
}

func getFibonacciSlice(fibonacciSeq *[]int32, x int32, y int32) []int32 {
	//TODO handle errors: x > y

	fibonacciSlice := (*fibonacciSeq)[x : y+1]
	return fibonacciSlice
}

func main() {
	listner, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFiboSlicerServer(grpcServer, &FiboSlicerServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalln(err)
	}
}
