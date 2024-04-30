package main

import "fmt"

//1 is the value of the neibor, numbers on the right is the number of elements(2)
//(1)                       2
//   \
//   (1 is the mid)
//    \
//   (2)                    3
//    \
//    (1 is the mid)
//     \
//     (3)                  4
//      \
//      (2 is the mid)

//mid = head
//if number of all elements is %2 than
//  mid += 1
//else do nothing

type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
  var mid *ListNode;
  mid = head;
  all_el_num := 1;
  for true{
    if head.Next != nil{
      all_el_num +=1;
      head = head.Next
      if all_el_num %2 ==0 {
        mid = mid.Next
      }
    }else{
      break
    } 
  }
  return mid
}

//a really smart way to do it by someone else
//func middleNode(head *ListNode) *ListNode {
//    if head == nil || head.Next == nil {
//        return head
//    }
//    slow, fast := head, head
//
//    for fast != nil && fast.Next != nil {
//        slow = slow.Next
//        fast = fast.Next.Next
//    }
//    return slow
//}

func main(){
  
  test := ListNode{1,&ListNode{2,&ListNode{3,nil}}} 
  fmt.Println(middleNode(&test).Val)

}
