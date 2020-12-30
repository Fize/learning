package marge_two_sorted_list

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
		Val:  9,
		Next: num2d,
	}
	num2d = &utils.ListNode{
		Val:  10,
		Next: nil,
	}

	input1 = num1a
	input2 = num2a
	output = "1, 1, 2, 4, 4, 5, 9, 10"
)

func TestMargeTwoSortedList(t *testing.T) {
	t.Logf("list1: %v, list2: %v, output: %v", utils.ListNodeToArray(input1), utils.ListNodeToArray(input2), output)
	r := MargeTwoSortedList(input1, input2)
	t.Logf("result: %v", utils.ListNodeToArray(r))
}
