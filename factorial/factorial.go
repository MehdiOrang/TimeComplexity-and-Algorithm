package main

import (
	"fmt"
	"time"
)

func factorial(n uint64) uint64 {
	var result uint64 = 1
	if n <= 1 {
		return 1
	}
	var i uint64 = 1
	for ; i <= n; i++ {
		result *= i
	}
	return result

}
func main() {
	start := time.Now()
	number := factorial(20)
	fmt.Println(number)
	t := time.Now()
	
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}

