package leetcode

import "log"

func getKthMagicNumber(k int) int {
	index3 := 0
	index5 := 0
	index7 := 0
	numList := make([]int, k)
	numList[0] = 1
	for i := 1; i < k; i++ {
		numList[i] = min(min(numList[index3]*3, numList[index5]*5), numList[index7]*7)
		log.Println(numList[i], index3, index5, index7)
		if numList[i] == numList[index3]*3 {
			index3++
		}
		if numList[i] == numList[index5]*5 {
			index5++
		}
		if numList[i] == numList[index7]*7 {
			index7++
		}
	}
	return numList[k-1]
}

/* func min(i, j int) int {
	if i > j {
		return j
	}
	return i
} */
