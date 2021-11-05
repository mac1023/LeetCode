package main

//链表的结构
type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {

}

func reverselist(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
