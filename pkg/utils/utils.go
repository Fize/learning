package utils

type ListNode struct {
	Prev *ListNode
	Val  int
	Next *ListNode
}

func ListNodeToArray(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
