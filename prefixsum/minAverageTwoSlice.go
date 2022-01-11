package main

import "fmt"

//Big (O**2)
func Solution(A []int) int {
	// write your code in Go 1.4
	lenA := len(A)
	var sum float32
	var avgBase float32 = float32(float32((A[0] + A[1])) / float32(2))
	var avg float32 = float32(float32((A[0] + A[1])) / float32(2))
	var min int = 0
	for i := 0; i < lenA; i++ {
		for j := i + 1; j < lenA; j++ {
			diff := float32(j - i)
			if i < lenA-1 {
				if i == j-1 {
					sum = float32((A[i] + A[j]))
					avg = float32(sum / float32(diff+1))
				} else {
					sum += float32(A[j])
					avg = float32(sum / float32(diff+1))
				}
				if avg < avgBase {
					avgBase = avg
					min = i
				}
			}
		}
	}

	return min
}

func main() {
	fmt.Println(Solution([]int{4, 2, 2, 5, 1, 5, 8}))
}

//https://app.codility.com/demo/results/trainingMGGPEH-HAX/
