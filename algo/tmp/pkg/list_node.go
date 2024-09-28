package pkg

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{val, nil}
}

func ArrayToLinkedList(nums []int) *ListNode {
	dummy := NewListNode(0)
	node := dummy
	for _, num := range nums {
		node.Next = NewListNode(num)
		node = node.Next
	}
	return dummy.Next
}
