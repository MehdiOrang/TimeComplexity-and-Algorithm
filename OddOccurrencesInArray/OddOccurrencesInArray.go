package main
import "fmt"

func Solution(A []int) int {
	len := len(A)
	result := 0
	for i := 0; i < len; i++ {
		result ^= A[i]
	}
	return result
}


func main(){
	fmt.Println(Solution([]int{1,2,4,2,1}) )
}

//https://app.codility.com/demo/results/trainingMWFKYD-X7H/
