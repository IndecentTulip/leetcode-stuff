package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {

  stack := []string{}
  result := []string{}
  var backtrack func(int, int)

  backtrack = func(open int, closed int)  {
    //fmt.Println("open", open, "closed", closed, "stack", stack, "result", result)

    if open == closed && closed == n {
      result = append(result, strings.Join(stack, ""))
      return
    }
    if open < n {
      stack = append(stack, "(")
      backtrack(open+1, closed)
      stack = stack[:len(stack)-1]
    }
    if closed < open{
      stack = append(stack, ")")
      backtrack(open, closed+1)
      stack = stack[:len(stack)-1]
    }

  }
  backtrack(0,0)
  return result
}


func main(){

  fmt.Println(generateParenthesis(3))
}
