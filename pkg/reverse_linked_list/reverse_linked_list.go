/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/
package reverse_linked_list

import "github.com/Fize/learning/pkg/utils"

func ReverseList(head *utils.ListNode) *utils.ListNode {
	var newHead *utils.ListNode
	var next *utils.ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}
