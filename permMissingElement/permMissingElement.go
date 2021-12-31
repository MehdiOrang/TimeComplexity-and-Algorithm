package main

import (
	"fmt"
	"time"
)

// Big O(N * log(N))
func Solution(A []int) int {
    start := time.Now()
    
    missedNumber := 0
    len          := len(A)
    sorted       := make([]int, len+1)

    if len == 0 { return 1 }

    for i := 0; i < len; i++{
        sorted[A[i]-1] = 1
    }

    for i := 0; i <= len; i++{
        if sorted[i] == 0 {
            missedNumber = i+1
        } 
    }

    t := time.Now()
    elapsed := t.Sub(start)

    fmt.Println(elapsed)

    return missedNumber
}

func main(){
  fmt.Println(Solution([]int{1,2,5,3}))
}

// https://app.codility.com/demo/results/trainingZ79XHH-5DE/
