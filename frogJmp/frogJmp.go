package main

import "fmt"
import "time"

// solution 1 Big O(1)
func Solution(X int, Y int, D int) int {
    start := time.Now()
    jumps := 0
    
    switch {
    case Y-X < D && Y > X:
         jumps = 1
    case Y-X < D :
        jumps = 0
    case Y%D > X:
        jumps = (Y/D)+((X+(Y%D))/D)
    default:
    	jumps = (Y/D)-((X-(Y%D))/D)
    }

    t := time.Now()
    elapsed := t.Sub(start)
    
    fmt.Println(elapsed)

    return jumps
}


/* solution 2 Big O(1)
func Solution(X int, Y int, D int) int {
    start := time.Now()

    jumps := 0
    if Y-X < D && Y > X {
        jumps = 1
    }else if Y-X < D {
        jumps = 0
    }else if Y%D > X {
       jumps  = (Y/D)+((X+(Y%D))/D)
    }else{
        jumps = (Y/D)-((X-(Y%D))/D)
    }
    
    t := time.Now()
    elapsed := t.Sub(start)
    
    fmt.Println(elapsed)

    return jumps
}
*/


/* solution 2, Big O(Y-X)
func Solution(X int, Y int, D int) int {
    start := time.Now()
    // write your code in Go 1.4
    jumps := 0

    for i := X; i < Y; i += D{
        jumps++
    }

    t := time.Now()
    elapsed := t.Sub(start)
    
    fmt.Println(elapsed)
    
    return jumps
}
*/


func main(){
	fmt.Println(Solution(5, 100000000000, 99))
}

//https://app.codility.com/demo/results/training4U3Q8N-YGV/
