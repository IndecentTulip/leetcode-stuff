package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {

  biggest := 0;
  for _,half := range accounts {
    sum := 0;
    for _,num := range half{
      sum += num;
    }
    if sum > biggest{
      biggest = sum;
    }
  }
  return biggest;
}

func main(){

  bank_info := [][]int{{1,2,3},{3,2,1}};

  fmt.Println(maximumWealth(bank_info));

  fmt.Println(maximumWealth_concurrency(bank_info))

}
