/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
*/
package palindrome_linked_list

import "github.com/Fize/learning/pkg/utils"

func IsPalindrome(head *utils.ListNode) bool {
	if head == nil {
		return true
	}
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	var start, end = 0, len(arr) - 1
	for {
		if start >= end {
			return true
		}
		if arr[start] != arr[end] {
			return false
		}
		start++
		end--
	}
}
