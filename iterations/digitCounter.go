package main

import (
	"fmt"
	"strconv"
)

func solution(A int) int {
	/*
		or
		counter := 0
		for A > 0 {
			A = A / 10
			counter++
		}
		return counter
	*/
	return len(strconv.Itoa(A))
}

func main() {
	fmt.Println(solution(9))
	fmt.Println(solution(299))
}
