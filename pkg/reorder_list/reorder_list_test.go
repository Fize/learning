package reorder_list

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

	input1 = num1a
	output = "1, 5, 2, 4"
)

func TestReorderList(t *testing.T) {
	t.Logf("list1: %v, except: %v", utils.ListNodeToArray(input1), output)
	r := ReorderList(input1)
	t.Logf("result: %v", utils.ListNodeToArray(r))
}
