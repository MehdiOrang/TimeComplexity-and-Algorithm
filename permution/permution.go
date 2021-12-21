package main

func Solution(A []int) int {
    // write your code in Go 1.4
    reflection := make([]int,len(A))
    for _, v := range A {
         if(0 <= v  && v > len(A) ){
             return 0
         }
        if reflection[v-1] == 1 {
            return 0
        }
        reflection[v-1] = 1
    }
    return 1
}

func main(){
  Solution([]int{1,2,3})
}
