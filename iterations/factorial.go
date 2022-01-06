package main

import "fmt"

func solution(A int) int {
	facatorial := 1

	for i := 1; i <= A; i++ {
		facatorial *= i
	}

	return facatorial
}

func main() {
	fmt.Println(solution(10))
}
