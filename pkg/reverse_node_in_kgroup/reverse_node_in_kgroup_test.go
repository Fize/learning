package reverse_node_in_kgroup

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
		Next: num1e,
	}
	num1e = &utils.ListNode{
		Val:  2,
		Next: num1f,
	}
	num1f = &utils.ListNode{
		Val:  1,
		Next: num1g,
	}
	num1g = &utils.ListNode{
		Val:  8,
		Next: nil,
	}

	input1 = num1a
	k = 3
	except = "4, 2, 1, 1, 2, 5, 8"
)

func TestReverseNodeInKGroup(t *testing.T) {
	t.Logf("list1: %v, except: %v", utils.ListNodeToArray(input1), except)
	r := ReverseNodeInKGroup(input1, k)
	t.Logf("result: %v", utils.ListNodeToArray(r))
}
