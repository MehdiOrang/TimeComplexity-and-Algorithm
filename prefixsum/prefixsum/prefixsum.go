package prefixsum

func PrefixSums(A []int) []int {
	length := len(A)
	prefixSlice := make([]int, length+1)
	for i := 1; i <= length; i++ {
		prefixSlice[i] = prefixSlice[i-1] + A[i-1]
	}
	return prefixSlice
}


