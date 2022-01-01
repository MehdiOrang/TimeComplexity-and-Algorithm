package main

import (
	"fmt"
	"time"
)

// Big O(N * log(N))
func Solution(A []int) int {
    start := time.Now()
    len  := len(A)
    head := A[0]
    tail := 0
    for i := 1; i < len; i++{
        tail += A[i]
    }
 
    min_diff := abs(head - tail)
    if len == 2{
        return min_diff
    }
    for i := 1; i < len; i++{
        head += A[i]
        tail -= A[i]
        if abs(head-tail) < min_diff{
            min_diff = abs(head - tail)
        }
    }
    
    t := time.Now()
    elapsed := t.Sub(start)

    fmt.Println(elapsed)

    return min_diff
}


func abs(N int) int { 
    if N >= 0{
        return N
    }else{
        return -N
    }
}

func main(){
  fmt.Println(Solution([]int{1,2,5,3}))
}

// https://app.codility.com/demo/results/trainingHDQXY6-BJB/



