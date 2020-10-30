package utils

import "fmt"

// PrintArray
func PrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("")
}
