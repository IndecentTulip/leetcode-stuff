package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {

  var result [][]int;
  var stack []int
  var helper func(int,int, *[][]int, []int)
  helper = func (i int, target int, result *[][]int, stack []int) {
    if (i == len(candidates)){
      fmt.Println("i=l, tar",target)
      fmt.Println("i=l, stack",stack)
      if target == 0{
        temp := make([]int, len(stack))
        copy(temp, stack)
        *result = append(*result, temp)
      }
      return
    }

    if candidates[i] <= target{
      stack = append(stack, candidates[i])
      fmt.Println("new tar",target - candidates[i])
      helper(i, target - candidates[i], result, stack)
      stack = stack[:len(stack)-1]
    }
    helper(i+1, target, result, stack)

  }

  helper( 0, target, &result, stack)

  return result
}

func main(){

  arr1 := []int{2,3,5}

  fmt.Println(combinationSum(arr1, 8))

}
