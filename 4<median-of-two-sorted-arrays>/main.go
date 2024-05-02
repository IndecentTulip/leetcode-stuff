package main

import "fmt"



func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  type node struct {
    val int
    left *node
    right *node
  }

  newNode := func(val int) *node {
      return &node{
          val:   val,
          left:  nil,
          right: nil,
      }
  }

  var insert func(*node, int) *node
  insert = func(root *node, val int) *node {
    if root == nil{
      return newNode(val)
    }

    if val == root.val{
      root.left = insert(root.left, val)
    }
    if val < root.val{
      root.left = insert(root.left, val)
    } 
    if val > root.val{
      root.right = insert(root.right, val)
    } 

    return root
  }

  var root *node
  for _,num := range nums1{
    root = insert(root, num)
  }
  for _,num := range nums2{
    root = insert(root, num)
  }

  result := []int{}
  var inorderSearch func(*node, []int) []int
  inorderSearch = func(root *node, result []int) []int{
    if root == nil{
      return nil
    }
    
    if root.left != nil {
      result = inorderSearch(root.left, result)
    }
    result = append(result, root.val)
    if root.right != nil {
      result = inorderSearch(root.right, result)
    }

    return result

  }

  result = inorderSearch(root, result)

  length := len(result)


  if !(length %2 ==0){
    return float64( result[length/2] )
  }
  if length %2 ==0{
    num := (float64(result[(length/2)-1]) + float64(result[length/2]))
    return float64(num/2)
  }

  return 0
}

// the correct way
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	merged := make([]int, 0)
//
//	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
//		if i < len(nums1) && (j == len(nums2) || nums1[i] <= nums2[j]) {
//			merged = append(merged, nums1[i])
//			i += 1
//		} else {
//			merged = append(merged, nums2[j])
//			j += 1
//		}
//
//	}
//
//	var median float64
//
//	mergedLen := len(merged)
//	if mergedLen%2 == 0 {
//		median = float64(merged[mergedLen/2]+merged[mergedLen/2-1]) / 2
//	} else {
//		median = float64(merged[mergedLen/2])
//	}
//
//	return median
//}

func main(){
  arr1 := []int{1,1}
  arr2 := []int{1,2}
  fmt.Println(findMedianSortedArrays(arr1, arr2))
}
