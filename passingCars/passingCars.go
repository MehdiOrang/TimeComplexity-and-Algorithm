package main

import "fmt"

func Solution(A []int) int {
	counter := 0
	incrementVal := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			incrementVal++
		}
		if A[i] == 1 {
			counter = counter + incrementVal
		}
		if counter > 1000000000 {
			return -1
		}
	}

	return counter
}

func main() {
	fmt.Println(Solution([]int{0, 1, 0, 1, 1}))
}

//https://app.codility.com/demo/results/trainingRZSGPY-UJQ/
