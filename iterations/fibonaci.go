package main

import (
	"fmt"
)

func solution(A int) []int {
	a := 0
	b := 1
	fibo := make([]int, 0)
	for a <= A {
		fibo = append(fibo, a)
		c := a + b
		a = b
		b = c
	}

	return fibo
}

func main() {
	fmt.Println(solution(20))
}
