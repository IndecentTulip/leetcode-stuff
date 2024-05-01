package main

import "fmt"

func searchInsert(nums []int, target int) int {
  right := len(nums) -1
  left := 0
  mid := 0 
  for left <= right{
    mid = left + (right - left) /2;

    if nums[mid] == target{
      return mid
    }

    if nums[mid] < target{
      left = mid + 1
    }else{
      right = mid - 1
    }
  }

  if left == mid && left == right+1{
    return mid
  }
  if left == right && right == mid{
    return mid
  }
  if right == mid && left-1 == right{
    return mid+1
  }
  return 0 
}

func main(){

  in := 2
  fmt.Println(in)
  nums := []int{3,5,7,9,10}
  fmt.Println(nums)
  fmt.Println(searchInsert(nums, in))

}
