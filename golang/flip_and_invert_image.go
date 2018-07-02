package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println(flipAndInvertImage([][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}}))
	fmt.Println(1 ^ 1)
	fmt.Println(0 ^ 1)
}

func flipAndInvertImage(A [][]int) [][]int {
	if len(A) <= 0 {
		return A
	}
	if len(A) != len(A[0]) {
		return A
	}
	tmp := []int{}
	result := [][]int{}
	for _, value := range A {
		for _, value1 := range value {
			if value1 < 0 || value1 > 1 {
				return [][]int{}
			}
			tmp = append([]int{value1 ^ 1}, tmp...)
		}
		result = append(result, tmp)
		tmp = []int{}
	}
	return result
}
