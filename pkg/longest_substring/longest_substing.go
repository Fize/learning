/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
package longest_substring

func LengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	current := make(map[byte]int)
	ans := 0
	for right < len(s) {
		if _, ok := current[s[right]]; ok {
			left++
			right = left
			current = map[byte]int{}
		} else {
			current[s[right]]++
			mid := right - left + 1
			if ans < mid {
				ans = mid
			}
			right++
		}
	}
	return ans
}
