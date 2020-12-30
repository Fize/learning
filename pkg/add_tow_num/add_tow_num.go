/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

package add_tow_num

import (
	"github.com/Fize/learning/pkg/utils"
)

func AddTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	resultHead := new(utils.ListNode)
	head := resultHead
	var sum int
	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		resultHead.Val = sum % 10
		sum = sum / 10
		if l1 != nil || l2 != nil || sum > 0 {
			resultHead.Next = new(utils.ListNode)
			resultHead = resultHead.Next
		}
	}
	return head
}
