package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	end := len(nums) - 1
	start := 0
	for (end - start) > 1 {
		mid := start + (end-start)/2
		if target < nums[mid] {
			end = mid
		}
		if target == nums[mid] {
			fmt.Println("index is:", mid)
			return
		}
		if target > nums[mid] {
			start = mid
		}
	}
	fmt.Println(-1)
}
