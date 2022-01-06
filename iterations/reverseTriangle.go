package main

import "fmt"

func solution(A int) {
	space := ""
	for i := A; i >= 1; i-- {
		space += " "
		for j := 1; j <= i; j++ {
			fmt.Printf(string(42) + " ")
		}
		fmt.Println()
		fmt.Printf(space)
	}
}

func main() {
	solution(4)
}
