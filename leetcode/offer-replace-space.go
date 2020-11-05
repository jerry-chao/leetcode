package leetcode

func replaceSpace(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			result = result + "%20"
		} else {
			result = result + string(s[i])
		}
	}
	return result
}
