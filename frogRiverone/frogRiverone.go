package main
import "fmt"
func Solution(X int, A []int) int {
    // write your code in Go 1.4
    result := -1
    max := 0
    reflection := make([]int, X)
    if(X < 1){
        return result
    }
    for i , v := range A {
        if(reflection[v-1] == 0){
            reflection[v-1] = 1
            max += 1
        }
        if(reflection[X-1] == 1 && max == X){
            result = i 
            break
        }
    }
    
    return result 
}

func main(){
	fmt.Println(Solution(5,[]int{1,1,1,2,3,4,2,5}))
}
