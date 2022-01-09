package main

import "fmt"

//Big O(N^2)
func BigON2(N int) int {
	result := 0
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			result += 1
		}
	}
	return result
}

//Big O(N + M)

//Big O(N)
func BigON(N int) int {
	result := 0
	for i := 0; i < N; i++ {
		result += i + 1
	}
	return result
}

//Big O(N)
func BigOlogN(N int) int {
	result := 0
	i := 1
	for k := N; k > (N / 2); k = k - 1 {
		result += i + k
		i++
	}
	return result
}

//Big O(1)
func BigO1(N int) int {
	result := N * (N + 1) / 2
	return result
}

func main() {
	N := 100000
	fmt.Println(BigON2(N))
	fmt.Println(BigON(N))
	fmt.Println(BigOlogN(N))
	fmt.Println(BigO1(N))
}
