package main
import "fmt"
/*
func Solution(N int, A []int) []int {
    // write your code in Go 1.4

    result := make([]int,N)
    max := 0
    for i:= 0; i <len(A); i++{
         if(A[i]<= N){
            result[A[i]-1] += 1
            if(result[A[i]-1] > max){
                max = result[A[i]-1] 
            }
         }
    }
    
    for j:= 0; j < N; j++{
                  result[j] = max
     }
    return result
}



*/

func Solution(N int, A []int) []int {
    // write your code in Go 1.4
    result := make([]int,N)
    max    := 0
    update := 0
    for i:=0; i<len(A); i++{
        if(A[i]<=N){
            if(result[A[i]-1] < update){
                result[A[i]-1] = update
            }
            result[A[i]-1] += 1
            if(result[A[i]-1] > max){
                max = result[A[i]-1] 
            }
        }else{
            update = max
        }
    }

    for j:=0; j< N; j++{
       if(result[j] < update){
           result[j] = update
       }
    } 
    return result
}


func main() {
	fmt.Println(Solution(1,[]int{1}))
}

