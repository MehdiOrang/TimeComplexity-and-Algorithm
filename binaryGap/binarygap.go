package main
// https://app.codility.com/demo/results/trainingRSDAKX-JX3/
import "fmt"

func Solution(N int) int {
    binary := make([]int,0)
	max_temp := 0
	max := 0
    flag := 1
		for flag != 0 {
		binary = append([]int{N%2}, binary...)
			N /= 2
			if(N <= 3){
				binary = append([]int{N%2}, binary...)
				binary = append([]int{N/2}, binary...)
				flag = 0
				break
			}
		}
		
	   for i:= 0 ; i < len(binary)-1; i++{
	   		if (binary[i] == 0 && binary[i+1] == 1 && max_temp > max){
			   max = max_temp
			}
	   		if (binary[i] == 0 && binary[i+1] == 1){
			   max_temp = 0
			}
			if (binary[i] == 1 && binary[i+1] == 0){
				max_temp += 1
			}
			if (binary[i] == 0 && max_temp != 0){
				max_temp += 1
			}			
	   }
	 fmt.Println(binary)
    return max
}

func main(){
	fmt.Println(Solution(561892))
}
