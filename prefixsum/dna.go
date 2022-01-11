package main

import "fmt"

//Big O(N *M)
// func Solution(S string, P []int, Q []int) []int {

// 	lenP := len(P)
// 	result := make([]int, lenP)

// 	for j := 0; j < lenP; j++ {
// 		result[j] = Nucleotide(S[P[j]])
// 		for k := P[j]; k <= Q[j]; k++ {
// 			if Nucleotide(S[k]) < result[j] {
// 				result[j] = Nucleotide(S[k])
// 			}
// 		}
// 	}

// 	return result
// }

// func Nucleotide(char byte) int {
// 	result := 1
// 	switch char {
// 	case 'A':
// 		result = 1
// 	case 'C':
// 		result = 2
// 	case 'G':
// 		result = 3
// 	case 'T':
// 		result = 4
// 	}
// 	return result
// }

// Big O(N+M)
func Solution(S string, P []int, Q []int) []int {
	// write your code in Go 1.4
	values := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	plate := [4]int{0, 0, 0, 0}
	counts := [][4]int{plate}
	result := []int{}

	for _, i := range S {
		plate[values[i]-1] += 1
		counts = append(counts, plate)

	}

	for i := 0; i < len(P); i++ {
		countsP := counts[P[i]]
		countsQ := counts[Q[i]+1]
		if Q[i] == P[i] {
			result = append(result, values[rune(S[Q[i]])])
		} else if countsQ[0] > countsP[0] {
			result = append(result, 1)
		} else if countsQ[1] > countsP[1] {
			result = append(result, 2)
		} else if countsQ[2] > countsP[2] {
			result = append(result, 3)
		} else if countsQ[3] > countsP[3] {
			result = append(result, 4)
		}
	}
	return result
}

func main() {
	fmt.Println(Solution("CCGAT", []int{1, 2, 1}, []int{3, 2, 1}))
}

//https://app.codility.com/demo/results/trainingRDPJBN-2RB/
