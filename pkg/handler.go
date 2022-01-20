package main

import "fmt"

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

func main() {
	fibonacciSequence := getFibonacciSequence(25)
	fmt.Println(getFibonacciSlice(fibonacciSequence, 3, 5))
}
