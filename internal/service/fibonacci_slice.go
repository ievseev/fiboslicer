package service

import fiboslicer "github.com/ievseev/fibonacci-slicer/internal"

type FibonacciSliceService struct {
}

func NewFibonacciSliceService() *FibonacciSliceService {
	return &FibonacciSliceService{}
}

func (s *FibonacciSliceService) GetFibonacciSlice(sequenceLimit fiboslicer.SequenceLimit) (*[]int, error) {
	sequenceLen := sequenceLimit.Y + 1
	fibonacciSequence := getFibonacciSequence(sequenceLen)
	result := getFibonacciSlice(fibonacciSequence, sequenceLimit.X, sequenceLimit.Y)
	return result, nil
}

func getFibonacciSequence(n int) *[]int {

	//TODO handle error: n<1

	if n == 1 {
		return &[]int{1}
	}

	fibonacciSeq := []int{1, 1}

	if n == 2 {
		return &fibonacciSeq
	}

	for i := 2; i < n; i++ {
		fibonacciSeq = append(fibonacciSeq, fibonacciSeq[i-1]+fibonacciSeq[i-2])
	}

	return &fibonacciSeq
}

func getFibonacciSlice(fibonacciSeq *[]int, x int, y int) *[]int {
	//TODO handle errors: x > y

	fibonacciSlice := (*fibonacciSeq)[x : y+1]
	return &fibonacciSlice
}
