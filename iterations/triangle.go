package main

import "fmt"

func solution(A int) {

	for i := 1; i <= A; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(string(42) + " ")
		}
		fmt.Println()
	}
}

func main() {
	solution(4)
}
