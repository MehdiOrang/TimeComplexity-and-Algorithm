let binary = "001011111111111111111111111111111111111111111111111111111111111111111111110000000000000000000000000111111110000000001010011111111111111111111111111111111111110"

function solution(S){
  let number = parseInt(S, 2);
  let steps = 0
  while(number > 0){
    if (number % 2 == 0){
      number  /= 2
    }else{
     number -= 1
    }
    steps += 1

   }
	return steps
}
