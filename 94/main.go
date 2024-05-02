package main

import "fmt"


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
// 14
func helper(root *TreeNode, result []int) []int{
  // 40 - root
   if root == nil{
    return nil
  }

  // 4 - root
  if root.Left != nil{
    result = helper(root.Left, result)
  }
  ///
  result = append(result, root.Val)
  if root.Right != nil{
    result = helper(root.Right, result)
  }

  return result 
} 

func inorderTraversal(root *TreeNode) []int {
  out := []int{};
  return helper(root, out)
}
func main(){


  var node3 *TreeNode = &TreeNode{3,nil,nil}
  node2 :=&TreeNode{2,node3,nil}
  node1 :=&TreeNode{1,nil,node2}
  fmt.Println(inorderTraversal(node1))
}
