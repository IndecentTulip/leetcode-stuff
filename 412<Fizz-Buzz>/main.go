package main

import "fmt"

func fizzBuzz(n int) []string {

  output := make([]string, n);
  for i:=0; i<n; i++{
    if (i+1)%3 == 0 && (i+1)%5 == 0{
      output[i] = "FizzBuzz";
    }else if (i+1)%3 == 0 {
      output[i] = "Fizz";
    }else if (i+1)%5 == 0 {
      output[i] = "Buzz";
    }else{
      output[i] = fmt.Sprintf("%v", i+1)
    }

  }
  return output
}

func main(){

  fmt.Printf("%s", fizzBuzz(15))
}
