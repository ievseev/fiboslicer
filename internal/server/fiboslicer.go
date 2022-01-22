package server

import (
	"context"
	pb "fiboslicer/pkg/github.com/ievseev/fiboslicer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FiboSlicerServer struct {
	pb.UnimplementedFiboSlicerServer
}

func (s *FiboSlicerServer) GetFibonacciSlice(ctx context.Context, req *pb.SliceRequest) (*pb.FibonacciSliceResponse, error) {
	if req.X > req.Y {
		err := status.Error(codes.InvalidArgument, "X must be less than Y")
		return nil, err
	}

	sequenceLen := req.Y + 1
	fibonacciSequence := getFibonacciSequence(sequenceLen)
	result := getFibonacciSlice(fibonacciSequence, req.X, req.Y)

	return &pb.FibonacciSliceResponse{Members: result}, nil
}

func getFibonacciSequence(n int32) *[]int32 {

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

	fibonacciSlice := (*fibonacciSeq)[x : y+1]
	return fibonacciSlice
}
