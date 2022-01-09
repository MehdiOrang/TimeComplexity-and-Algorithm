package main

import "fmt"

func reverse(A []int) []int {
	len := len(A)
	reversed := make([]int, len)
	for i := 0; i <= (len / 2); i++ {
		reversed[i], reversed[len-i-1] = A[len-i-1], A[i]
	}
	return reversed
}

func main() {
	fmt.Println(reverse([]int{1, 2, 3, 4, 5, 6}))

}
