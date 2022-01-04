package main

import "fmt"

func Solution(A []int) int {
	pairs := 0
	counter := true
	len := len(A)
	for i := 0; i < len; i++ {
		for {
			if !counter {
				break
			}
		}
	}

	return pairs
}

func main() {
	fmt.Println(Solution([]int{1, 2, 4, 2, 1}))
}
