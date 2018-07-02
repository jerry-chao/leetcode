package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	list := []int{1, 2, 3}
	fmt.Println(list)
	fmt.Println(reverse(list))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	maxIndex, _ := maxInt(height)
	peakSum := sumWater(height[:maxIndex])
	topSum := sumWater(reverse(height[maxIndex+1:]))
	return peakSum + topSum
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func sumWater(height []int) int {
	peak, sum := 0, 0
	for _, value1 := range height {
		if peak > value1 {
			sum += peak - value1
		} else {
			peak = value1
		}
	}
	return sum
}

func maxInt(height []int) (int, int) {
	maxIndex, maxValue := 0, 0
	for index, value := range height {
		if value > maxValue {
			maxValue = value
			maxIndex = index
		}
	}
	return maxIndex, maxValue
}
