package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums0 := []int{2, 7, 11, 15}
	target0 := 9
	l1 := TwoSum(nums0, target0)
	expect0 := []int{0, 1}
	t.Logf("nums: %v, target: %d, expect: %v, result: %v", nums0, target0, expect0, l1)

	nums1 := []int{3, 5, 11, 11, 16}
	target1 := 22
	l2 := TwoSum(nums1, target1)
	expect1 := []int{2, 3}
	t.Logf("nums: %v, target: %d, expect: %v, result: %v", nums1, target1, expect1, l2)
}
