package service

import fiboslicer "github.com/ievseev/fibonacci-slicer/internal"

type FibonacciSlice interface {
	GetFibonacciSlice(sequenceLimit fiboslicer.SequenceLimit) (*[]int, error)
}

type Service struct {
	FibonacciSlice
}

func NewService() *Service {
	return &Service{FibonacciSlice: NewFibonacciSliceService()}
}
