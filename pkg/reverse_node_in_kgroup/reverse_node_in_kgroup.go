/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]

提示：

	列表中节点的数量在范围 sz 内
	1 <= sz <= 5000
	0 <= Node.val <= 1000
	1 <= k <= sz
*/

package reverse_node_in_kgroup

import (
	"github.com/Fize/learning/pkg/utils"
)

func reverseListNode(head *utils.ListNode, k int) (*utils.ListNode, *utils.ListNode) {
	var listNodeArray []*utils.ListNode
	for head != nil {
		listNodeArray = append(listNodeArray, head)
		head = head.Next
	}
	head = listNodeArray[0]
	if len(listNodeArray) < k {
		return head, listNodeArray[len(listNodeArray) - 1]
	}
	var newHead *utils.ListNode
	var next *utils.ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	
	return newHead, listNodeArray[0]
}

func ReverseNodeInKGroup(head *utils.ListNode, k int) *utils.ListNode {
	var listNodeArray []*utils.ListNode
	if head == nil {
		return head
	}
	for head != nil {
		listNodeArray = append(listNodeArray, head)
		head = head.Next
	}
	if len(listNodeArray) < k {
		return head
	}

	first, end := 0, k

	var newHead *utils.ListNode
	var newLast *utils.ListNode
	for {
		var newList []*utils.ListNode
		if first > len(listNodeArray) {
			break
		}
		if end > len(listNodeArray) {
			newList = listNodeArray[first:]
		} else {
			newList = listNodeArray[first:end]
		}
		if len(newList) <= 1 && end < len(newList) {
			newHead = listNodeArray[0]
			break
		}
		newList[len(newList) - 1].Next = nil
		head, last := reverseListNode(newList[0], k)
		if first == 0 {
			newHead = head
			newLast = last
		} else {
			newLast.Next = head
			newLast = last
		}
		first = first + k
		end = end + k
	}
	return newHead
}