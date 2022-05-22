package longest_palindromic_substring_test

import "testing"

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
// Related Topics 字符串 动态规划 👍 5062 👎 0

func TestV1(t *testing.T) {
	t.Log(longestPalindrome("aa"))
	t.Log(longestPalindrome("cbbd"))
	t.Log(longestPalindrome("babad"))
	t.Log(longestPalindrome("abcdefedcba"))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := 0
	start := 0
	for i := 0; i < len(s); i++ {
		// 奇数
		len1 := expandAroundCenter(s, i, i)
		// 偶数
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > maxLen {
			maxLen = len
			start = i - (len-1)/2
		}
	}
	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
