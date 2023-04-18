package main

// import (
// 	"fmt"
// )

// func main(){

// 	li_1 := ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 4,
// 				Next: nil,
// 			},
// 		},
// 	}

// 	li_2 := ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 3,
// 			Next: &ListNode{
// 				Val: 4,
// 				Next: nil,
// 			},
// 		},
// 	}

// 	fmt.Print(mergeTwoLists(&li_1, &li_2))

// }

// type ListNode struct {
//       Val int
//       Next *ListNode
// }

// // Soltuion
// // RunTime 4ms
// // Memory 2.9MB
// // func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

// // 	var arr []int
// // 	for list1 != nil {
// // 		arr = append(arr, list1.Val)
// // 		list1 = list1.Next
// // 	}

// // 	for list2 != nil{
// // 		arr = append(arr, list2.Val)
// // 		list2 = list2.Next
// // 	}

// // 	sort.Ints(arr)

// // 	dummy := new(ListNode)
// // 	node := dummy
// // 	for _, v := range arr {
// // 		list := ListNode{Val: v}
// // 		node.Next = &list
// // 		node = node.Next
// // 	}

// // 	return dummy.Next
// // }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummy := new(ListNode)
// 	n := dummy

// 	var f1, f2 = false, false
// 	for !f1 && !f2 {

// 		if list1.Val >=list2.Val {

// 			n.Next = list2
// 			list2 = list2.Next
// 		}else if list1.Val <= list2.Val {

// 			n.Next = list1
// 			list1 = list1.Next
// 		}else if list1 == nil  {

// 			n.Next = list2
// 			list2 = list2.Next
// 			f1 = true
// 		}else if list2 == nil{
// 			n.Next = list1
// 			list1 = list1.Next
// 			f2 = true
// 		}

// 		n = n.Next
// 	}

// 	printLi := dummy.Next
// 	for printLi != nil {
// 		fmt.Println(printLi.Val)
// 		printLi = printLi.Next
// 	}

// 	return dummy.Next
// }
