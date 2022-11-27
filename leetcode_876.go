package main

// import "fmt"

//  type ListNode struct {
//       Val int
//       Next *ListNode
// }

// func main(){
// 	h := middleNode(&ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val : 2,
// 			Next: &ListNode{
// 				Val : 3,
// 				Next:
// 				&ListNode{
// 					Val : 4,
// 					Next:
// 					&ListNode{
// 						Val : 5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	})

// 	fmt.Println(h)
// 	fmt.Println(h.Next)
// 	fmt.Println(h.Next.Next)
// 	fmt.Println(h.Next.Next.Next)

// }

// // Runtime 100%, Memory 14%
// func middleNode(head *ListNode) *ListNode {

// 	crnList := head
// 	i := 0
// 	for crnList != nil {
// 		i += 1
// 		crnList = crnList.Next
// 	}

// 	mid := int(i/2)
// 	j := 0
// 	for head!= nil{

// 		if j < mid {
// 			head = head.Next
// 			j++
// 		}else{
// 			break
// 		}

// 	}

// 	return head
// }