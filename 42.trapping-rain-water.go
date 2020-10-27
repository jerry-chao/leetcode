package main

import "fmt"

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
/* func trap(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	max := height[0]
	trapList := make([]int, n)
	stack := []int{}
	for i := 1; i < n; i++ {
		if height[i] >= max {
			// pop stack
			for len(stack) > 0 {
				topIndex := stack[len(stack)-1]
				trapList[topIndex] = max - height[topIndex]
				fmt.Println(topIndex, trapList[topIndex])
				stack = stack[:len(stack)-1]
			}
			max = height[i]
		} else {
			stack = append(stack, i)
		}
	}
	fmt.Println("stack len", max, len(stack))
	rightMax := 0
	for len(stack) > 0 {
		rightIndex := stack[len(stack)-1]
		if height[rightIndex] <= rightMax {
			trapList[rightIndex] = rightMax - height[rightIndex]
		} else {
			rightMax = height[rightIndex]
		}
		fmt.Println("rightMax:", rightMax)
		stack = stack[:len(stack)-1]
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + trapList[i]
	}
	return sum
} */

func trap(height []int) int {
	num := len(height)
	if num == 0 {
		return 0
	}

	leftMax := 0
	volume := 0
	stack := []int{}
	for i := 0; i < num; i++ {
		if height[i] >= leftMax {
			// pop all stack
			for len(stack) > 0 {
				last := stack[len(stack)-1]
				volume = volume + leftMax - height[last]
				fmt.Println(volume, leftMax, height[last])
				stack = stack[:len(stack)-1]
			}
			leftMax = height[i]
		} else {
			stack = append(stack, i)
		}
	}
	fmt.Println(volume, len(stack))
	// handle last stack
	rightMax := 0
	for len(stack) > 0 {
		lastValue := height[stack[len(stack)-1]]
		if lastValue < rightMax {
			volume = volume + rightMax - lastValue
		} else {
			rightMax = lastValue
		}
		stack = stack[:len(stack)-1]
	}
	return volume

}

// @lc code=end

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("trap:", trap(height), trap(height) == 6)
}
