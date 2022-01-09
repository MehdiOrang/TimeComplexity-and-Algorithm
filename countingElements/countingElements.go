package main

import (
	"fmt"
)

//improve with negative numbers Big O(N + M)
func solution(A []int, M int) []int {
	len := len(A)
	counts := make([]int, M+1)
	for i := 0; i < len; i++ {
		counts[A[i]] += 1
	}
	return counts
}

func main() {
	fmt.Println(solution([]int{1, 2, 4, 6, 2, 3, 2, 4, 3, 3}, 6))
}
