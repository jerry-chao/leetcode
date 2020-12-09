package leetcode

/*
 * @lc app=leetcode id=493 lang=golang
 *
 * [493] Reverse Pairs
 */

// @lc code=start

var countReversePairs int

func reversePairs(nums []int) int {
	mergeSort(nums, 0, len(nums)-1)
	PrintArray(nums)
	return countReversePairs
}

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)>>1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	mergeInternal(nums, left, mid, right)
}

func mergeInternal(nums []int, left, mid, right int) {
	tmp := []int{}
	first, second := left, mid+1
	for first <= mid && second <= right {
		if nums[first] > nums[second]*2 {
			countReversePairs++
		}
		if nums[first] < nums[second] {
			tmp = append(tmp, nums[first])
			first++
		} else {
			tmp = append(tmp, nums[second])
			second++
		}
	}
	for first <= mid {
		tmp = append(tmp, nums[first])
		first++
	}
	for second <= right {
		tmp = append(tmp, nums[second])
		second++
	}
	// set nums
	for i := 0; i < len(tmp); i++ {
		nums[left+i] = tmp[i]
	}
}

// @lc code=end
