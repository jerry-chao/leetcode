package leetcode

func getKthMagicNumber(k int) int {
	i3, i5, i7 := 0, 0, 0
	retList := make([]int, k)
	retList[0] = 1
	for i := 1; i < k; i++ {
		retList[i] = min(min(retList[i3]*3, retList[i5]*5), retList[i7]*7)
		if retList[i] == retList[i3]*3 {
			i3++
		}
		if retList[i] == retList[i5]*5 {
			i5++
		}
		if retList[i] == retList[i7]*7 {
			i7++
		}
	}
	return retList[k-1]
}

/* func min(i, j int) int {
	if i > j {
		return j
	}
	return i
} */
