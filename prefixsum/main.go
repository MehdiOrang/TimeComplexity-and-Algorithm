package main

import (
	"fmt"
	"prefixsum/prefixsum"
)

func main() {

	fmt.Println(prefixsum.PrefixSums([]int{2, 3, 7, 5, 1, 3, 9}))
	fmt.Println(mushroom([]int{2, 3, 7, 5, 1, 3, 9}, 4, 6))
}

func mushroom(A []int, k int, m int) int {
	len := len(A)
	result := 0
	prefixSum := prefixsum.PrefixSums(A)
	for i := 1; i < min(k, m)+1; i++ {
		leftPos := k - i
		rightPos := min(len-1, max(k, k+m-2*i))
		result = max(result, countTotal(prefixSum, leftPos, rightPos))
	}

	for i := 0; i < min(m+1, len-k); i++ {
		leftPos := k + i
		rightPos := max(0, min(k, k-(m-2*i)))
		result = max(result, countTotal(prefixSum, leftPos, rightPos))
	}

	return result
}

func min(A int, B int) int {
	if A > B {
		return B
	}

	return A
}

func max(A int, B int) int {
	if A > B {
		return A
	}

	return B
}

func countTotal(A []int, x int, y int) int {
	return A[y+1] - A[x]
}
