package main

import "fmt"

func main() {
	fmt.Println(countDiv(0, 1000000, 1))
}

func countDiv(A int, B int, K int) int {
	// write your code in Go 1.4

	result := 0

	if A == B && A%K == 0 {
		result = 1
	} else if A == B && A%K != 0 {
		result = 0
	} else if A%K == 0 {
		result = 1
	}

	result = (B / K) - (A / K) + result
	return result
}
