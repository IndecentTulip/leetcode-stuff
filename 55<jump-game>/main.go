package main

import (
	"fmt"
)

func canJump(nums []int) bool {

  //[1,4,5,0,4]
  //0 no  fin 4
  // 4 - 2 = 2
  // val of index > 2
  // fin index - num index = range
  //5 yes fin 2
  //4 yes fin 1
  //1 yes fin 0

  fin := len(nums)-1
  for i := len(nums) - 2; i >= 0; i-- {
    fmt.Println("step")
    if nums[i] >= fin - i  {
      fin = i
    }

  }

  if fin == 0{
    return true
  }
	return false
}




func main() {
	arr1 := []int{2,0,0,4}
  fmt.Println(arr1)
	fmt.Println(canJump(arr1))
}
