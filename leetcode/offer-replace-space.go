package leetcode

func replaceSpace(s string) string {
	return replaceSpaceRec(s, "")
}

func replaceSpaceRec(s string, t string) string {
	if len(s) > 0 {
		if string(s[0]) == " " {
			t = t + "%20"
		} else {
			t = t + string(s[0])
		}
		return replaceSpaceRec(s[1:], t)
	}
	return t
}
