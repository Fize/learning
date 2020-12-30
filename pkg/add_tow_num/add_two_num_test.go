package add_tow_num

import (
	"testing"

	"github.com/Fize/learning/pkg/utils"
)

var (
	num1a = &utils.ListNode{
		Val:  9,
		Next: num1b,
	}
	num1b = &utils.ListNode{
		Val:  9,
		Next: num1c,
	}
	num1c = &utils.ListNode{
		Val:  9,
		Next: num1d,
	}
	num1d = &utils.ListNode{
		Val:  9,
		Next: num1e,
	}
	num1e = &utils.ListNode{
		Val:  9,
		Next: num1f,
	}
	num1f = &utils.ListNode{
		Val:  9,
		Next: num1g,
	}
	num1g = &utils.ListNode{
		Val:  9,
		Next: nil,
	}

	num2a = &utils.ListNode{
		Val:  9,
		Next: num2b,
	}
	num2b = &utils.ListNode{
		Val:  9,
		Next: num2c,
	}
	num2c = &utils.ListNode{
		Val:  9,
		Next: num2d,
	}
	num2d = &utils.ListNode{
		Val:  9,
		Next: nil,
	}

	input1 = num1a
	input2 = num2a

	except = "[8 9 9 9 0 0 0 1]"
)

func TestAddTwoNum(t *testing.T) {
	t.Logf("input1: %v, intput2: %v, except: %s", utils.ListNodeToArray(input1), utils.ListNodeToArray(input2), except)
	output1 := AddTwoNumbers(input1, input2)
	t.Logf("result: %v", utils.ListNodeToArray(output1))
}
