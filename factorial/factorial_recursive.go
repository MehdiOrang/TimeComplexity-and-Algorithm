package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}


func main() {
	start := time.Now()
	
	number := factorial(20)
	fmt.Println(number)
	
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}

