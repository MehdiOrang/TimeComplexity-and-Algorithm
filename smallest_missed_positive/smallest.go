package main

import "fmt"

func Solution(A []int) int {
    // write your code in Go 1.4
    n := len(A)
    var present = make([]bool,n+1)

    for j:=0; j < n; j++{
        if(A[j] > 0 && A[j] <= n){
            present[A[j]] = true
        }
    }

    for z := 1; z <= n; z++{
           if (!present[z]){
                return z;
            }
    }

    return n + 1;
}

func main(){
	fmt.Println(Solution([]int{-1,-2,-3,-6,-5,-7,-4}))
	fmt.Println(Solution([]int{1,2,3,6,5,7,4}))
	fmt.Println(Solution([]int{1,2,3,6,-5,7,4}))
}

