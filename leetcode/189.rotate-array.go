package leetcode

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func Rotate(nums []int, k int) {
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	// rotateLoop(nums, k)
	rotateReverse(nums, k)
}

func rotateLoop(nums []int, k int) {
	n := len(nums)
	if n < 2 || k < 1 {
		return
	}
	k = k % n
	count := 0
	for i := 0; count < n; i++ {
		current := i
		step := 0
		pre := nums[current]
		for current != i || step == 0 {
			nextIndex := (current + k) % n
			tmp := nums[nextIndex]
			nums[nextIndex] = pre
			pre = tmp
			current = nextIndex
			count++
			step++
		}
	}
}

func rotateReverse(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	n := len(nums)
	k = k % n
	// reverse all array
	reverseArray(nums)
	// reverse first k
	reverseArray(nums[:k])
	// reverse last n -k
	reverseArray(nums[k:])
}

func reverseArray(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}
}

// @lc code=end
