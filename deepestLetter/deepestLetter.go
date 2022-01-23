package main

import (
    "fmt"
)

func getDeepestLetter(s string) rune {
  currentMax := 0
  mainMax := 0
  result := '?'
  length := len(s)
  for i := 0; i< length; i++{
    if s[i] == '('{
	currentMax++
	if currentMax > mainMax{
	  mainMax = currentMax
          result =rune(s[i+1])
	}
    }else if  s[i] == ')'{
	if currentMax > 0 {
	  currentMax--
	}else{
	  return '?'
	}
    }
  }

  if currentMax != 0 {
	return '?'
  }

  return result

}

func main(){
  fmt.Printf("%c\n",getDeepestLetter("a(b)c"))
  fmt.Printf("%c\n",getDeepestLetter("((A)(b)c"))
  fmt.Printf("%c\n",getDeepestLetter("(8)"))
  fmt.Printf("%c\n",getDeepestLetter("(!)"))
  fmt.Printf("%c\n",getDeepestLetter("((a))(((M)))(c)(D)(e)(((f))(((G))))h(i)"))
  fmt.Printf("%c\n",getDeepestLetter("( p((q)) ((s)t) )"))
}
