package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

// removeDuplicates ...
func removeDuplicates(nums []int) (int, []int) {
	if len(nums) < 2 {
		return len(nums), nums
	}
	l := 1
	pre := nums[0]
	for _, value := range nums {
		if pre != value {
			nums[l] = value
			pre = value
			l++
		}
	}
	return l, nums[:l]
}
