package main

import (
	"fmt"
	"prefixsum/prefixsum"
)

func main(){
	fmt.Println(prefixsum.PrefixSums([]int{2,3,7,5,1,3,9}))
}