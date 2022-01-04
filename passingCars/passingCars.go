package main

import "fmt"

func Solution(A []int) int {
	counter := 0
	incrementVal := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			incrementVal++
		}
		if A[i] == 1 {
			counter = counter + incrementVal
		}
		if counter > 1000000000 {
			return -1
		}
	}

	return counter
}

/*
func Solution(A []int) int {
    // write your code in Go 1.4
    pairs   := 0
    counter := 0
    len     := len(A)
    for i := 0; i <len; i++{
       // fmt.Println("print outer:", A[i])
        if A[i] == 0 {
            counter++
        }
        if (counter > 0){
            if (A[i] == 1){
                pairs = pairs + counter;
                if (pairs > 1000000000){
                        return -1;
                }
            }
        }
    }

    return pairs
}
*/
func main() {
	fmt.Println(Solution([]int{0, 1, 0, 1, 1}))
}

//https://app.codility.com/demo/results/trainingRZSGPY-UJQ/
