package test

func longestPalindromic(s string) string {
	if len(s) < 2 {
		return s
	}
	var maxLen int
	var start int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > maxLen {
			maxLen = len
			start = i - (len-1)/2
		}
	}
	return s[start : start+maxLen]
}

func max(len1 int, len2 int) int {
	if len1 > len2 {
		return len1
	}
	return len2
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

/*
    abc
   a: -1 0 1
   b:  0 1 2
*/
