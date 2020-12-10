package leetcode

/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start
func sortArray(nums []int) []int {
	// quickSort(nums, 0, len(nums)-1)
	mergeSortArray(nums, 0, len(nums)-1)
	return nums
}

func mergeSortArray(nums []int, l, r int) {
	if l < r {
		mid := l + (r-l)>>1
		mergeSortArray(nums, l, mid)
		mergeSortArray(nums, mid+1, r)
		mergeArray(nums, l, mid, r)
	}
}

func mergeArray(nums []int, l, mid, r int) {

	first, second := l, mid+1
	tmp := []int{}
	for first <= mid && second <= r {
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
	for second <= r {
		tmp = append(tmp, nums[second])
		second++
	}
	copy(nums[l:r+1], tmp)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	index := l - 1
	for j := l; j < r; j++ {
		if nums[j] < nums[r] {
			index++
			nums[index], nums[j] = nums[j], nums[index]
		}
	}
	index++
	nums[index], nums[r] = nums[r], nums[index]
	quickSort(nums, l, index-1)
	quickSort(nums, index+1, r)
}

// @lc code=end
