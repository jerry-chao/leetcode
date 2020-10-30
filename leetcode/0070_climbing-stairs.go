func climbStairs(n int) int {
	return climbStairsMap(n)
}

var status = map[int]int{1: 1, 2: 2}

func climbStairsMap(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	num, ok := status[n]
	if ok {
		return num
	} else {
		newNum := climbStairsMap(n-1) + climbStairsMap(n-2)
		status[n] = newNum
		return newNum
	}
}
