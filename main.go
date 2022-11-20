package main
 
 type ListNode struct {
      Val int
      Next *ListNode
}

func main(){
	reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val : 2,
			Next: &ListNode{
				Val : 3,
				Next: 
				&ListNode{
					Val : 4,
					Next: 
					&ListNode{
						Val : 5,
						Next: nil,
					},
				},
			},
		},
	})

 	// reverseList(&ListNode{})
}



func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 5
	r := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return r
}