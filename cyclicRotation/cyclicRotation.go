package main

import ( 
	"fmt"
	"math"
)

func Solution(A []int, K int) []int {

    len := len(A)
    if K == 0 || len == 0{
        return A
    }

    result := make([]int, len)
    for i := 0; i < len; i++{
        if i < (K%len){
            result[i] = A[(len - (K%len))+i];
        }else{
            result[i] = A[int(math.Abs(float64((K%len)-i)))];
        }
            
    }

    return result
}

func main(){
	fmt.Println(Solution([]int{2,3,4,5,6}, 3))
}

// https://app.codility.com/demo/results/trainingW3TXVC-VHP/
