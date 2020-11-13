package leetcode

func getKthMagicNumber(k int) int {
	i3, i5, i7 := 0, 0, 0
	numList := make([]int, k)
	numList[0] = 1
	for i := 1; i < k; i++ {
		numList[i] = min(min(numList[i3]*3, numList[i5]*5), numList[i7]*7)
		if numList[i] == numList[i3]*3 {
			i3++
		}
		if numList[i] == numList[i5]*5 {
			i5++
		}
		if numList[i] == numList[i7]*7 {
			i7++
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
