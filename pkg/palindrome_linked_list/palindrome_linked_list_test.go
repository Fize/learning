package palindrome_linked_list

import (
	"testing"

	"github.com/Fize/learning/pkg/utils"
)

var (
	num1a = &utils.ListNode{
		Val:  1,
		Next: num1b,
	}
	num1b = &utils.ListNode{
		Val:  2,
		Next: num1c,
	}
	num1c = &utils.ListNode{
		Val:  4,
		Next: num1d,
	}
	num1d = &utils.ListNode{
		Val:  5,
		Next: nil,
	}

	num2a = &utils.ListNode{
		Val:  1,
		Next: num2b,
	}
	num2b = &utils.ListNode{
		Val:  4,
		Next: num2c,
	}
	num2c = &utils.ListNode{
		Val:  4,
		Next: num2d,
	}
	num2d = &utils.ListNode{
		Val:  1,
		Next: nil,
	}

	input1  = num1a
	input2  = num2a
	output1 = "false"
	output2 = "true"
)

func TestPalindromeLinkedList(t *testing.T) {
	t.Logf("list1: %v, list2: %v, output1: %v, output2: %v\n", utils.ListNodeToArray(input1), utils.ListNodeToArray(input2), output1, output2)
	r1 := IsPalindrome(input1)
	r2 := IsPalindrome(input2)
	t.Logf("result1: %v, result2: %v\n", r1, r2)
}
