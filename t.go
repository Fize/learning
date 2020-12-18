package main

import "fmt"

func main() {
	s := "abcabckJ"
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
	fmt.Println(ans)
}
