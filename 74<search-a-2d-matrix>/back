package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {

  last :=0
  first :=0
  left := 0 //0
  right := len(matrix)-1 //4
  mid :=0
  for left <= right {
    mid = left + (right - left)/2 // 2 +(3-2)/2
    last = matrix[mid][len(matrix[0])-1]
    first = matrix[mid][0]

    if last == target || first == target{
      return true
    }
    if last < target && first < target{
      left = mid + 1
    }
    if last > target && first > target{
      right = mid - 1
    }
    if last > target && first < target{
      return binary_search(matrix[mid], target)
    }

  } 
  return binary_search(matrix[mid], target)
}

func binary_search(arr []int, target int ) bool{

  left := 0
  right := len(arr) -1
  mid :=0

  for left <= right{
    mid = left + (right- left)/2

    if arr[mid] < target {
      left = mid +1
    }
    if arr[mid] > target {
      right = mid -1
    }
    if arr[mid] == target{
      return true
    }
  } 

  return false
}

func main(){

  test2 := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,50}}


  fmt.Println(searchMatrix(test2, 1))
  fmt.Println(searchMatrix(test2, 3))
  fmt.Println(searchMatrix(test2, 5))
  fmt.Println(searchMatrix(test2, 7))
  fmt.Println(searchMatrix(test2, 10))
  fmt.Println(searchMatrix(test2, 11))
  fmt.Println(searchMatrix(test2, 16))
  fmt.Println(searchMatrix(test2, 20))
  fmt.Println(searchMatrix(test2, 23))
  fmt.Println(searchMatrix(test2, 30))
  fmt.Println(searchMatrix(test2, 34))
  fmt.Println(searchMatrix(test2, 50))

}
