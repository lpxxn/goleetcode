package test

import "testing"

func longestPalindromicSubstring(s string) string {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func expandAroundCenter(s string, l int, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func TestName(t *testing.T) {
	t.Log(longestPalindromicSubstring("babad"))
	t.Log(longestPalindromicSubstring("cbbd"))
	t.Log(longestPalindromicSubstring("abcdefedcba"))
}

/*
    abc
   a: -1 0 1
   b:  0 1 2
*/
