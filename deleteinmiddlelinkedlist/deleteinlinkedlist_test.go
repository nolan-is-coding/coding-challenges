package deleteinmiddlelnkedlist

import (
	"reflect"
	"testing"
)

func storeList(head *ListNode) []int {
	tmp := head
	arr := []int{}
	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
	return arr
}

func TestWiggleSort(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{2, nil}
	n3 := ListNode{3, nil}
	n4 := ListNode{4, nil}
	n5 := ListNode{5, nil}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	l := removeMiddleLinkedList(&n1)
	arr := storeList(l)
	expected := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(arr, expected) {
		t.Fatal("not good")
	}
}
