package deleteinmiddlelnkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func listLen(head *ListNode) int {
	tmp := head
	i := 0
	for tmp != nil {
		i += 1
		tmp = tmp.Next
	}
	return i
}

// 1 -> 2 -> 4 -> 5
// 1 -> 2 -> 4 -> Nil 
// count the elem / 2 -> middle 
func removeMiddleLinkedList(head *ListNode) *ListNode {
	l := listLen(head)

	if l <= 1 {
		return nil
	} 
	tmp := head
	i := 0
	for i < (l/2) -1{
		tmp=tmp.Next
		i++
	}
	if tmp.Next != nil && tmp.Next.Next != nil {
		tmp.Next = tmp.Next.Next
	} else if tmp.Next == nil || tmp.Next.Next == nil {
		tmp.Next = nil
	}
	return head
}

func main() {
	fmt.Println("nothing")
}
