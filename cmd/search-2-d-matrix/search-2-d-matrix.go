package main

import (
	"fmt"
	"leetcode/leetcode"
)

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(leetcode.SearchMatrix(matrix, 3))
}
