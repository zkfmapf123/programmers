package main

// import (
// 	"fmt"
// )

// type ListNode struct {
// 	Val int
// 	Next *ListNode
// }

// func main(){
// 	li := ListNode{
// 		Val: 2,
// 		Next : &ListNode{
// 			Val : 4,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: nil,
// 			},
// 		},
// 	}

// 	li2 := ListNode{
// 		Val: 5,
// 		Next : &ListNode{
// 			Val : 6,
// 			Next: &ListNode{
// 				Val: 4,
// 				Next: nil,
// 			},
// 		},
// 	}

// 	fmt.Println((addTwoNumbers(&li,&li2)))
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	head := new(ListNode)
// 	cur := head

// 	isOver := false

// 	for true  {
// 		n1, n2 := 0,0

// 		if l1 != nil {
// 			n1 = l1.Val
// 			l1 = l1.Next
// 		}

// 		if l2 != nil {
// 			n2 = l2.Val
// 			l2 = l2.Next
// 		}

// 		sum := n1 + n2
// 		if isOver {
// 			sum += 1
// 			isOver = false
// 		}

// 		if sum >= 10 {
// 			sum -= 10
// 			isOver = true
// 		}

// 		cur.Next = &ListNode{Val : sum}
// 		cur = cur.Next

// 		if l1 == nil && l2 == nil && !isOver {
// 			break
// 		}
// 	}

// 	return head.Next
// }
