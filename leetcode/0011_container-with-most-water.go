func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	p := 0
	q := len(height) - 1
	maxArea := 0
	for q > p {
		newArea := area(p, q, height[p], height[q])
		maxArea, _ = sort(maxArea, newArea)
		if height[p] > height[q] {
			q--
		} else {
			p++
		}
	}
	return maxArea
}

func area(i, j, hi, hj int) int {
	_, min := sort(i, j)
	return min * (j - i)
}

func sort(i, j int) (int, int) {
	if i > j {
		return i, j
	}
	return j, i
}
