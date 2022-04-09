package test

import (
	"fmt"
	"testing"
)

/*
Given a string, find the length of the longest substring without repeating characters.

**Example:**
```
Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```
*/

// Given a string, find the length of the longest substring without repeating characters

func TestStr1(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabbzycdebfdb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("aaa"))
	fmt.Println(lengthOfLongestSubstring("aaaa"))
	fmt.Println(lengthOfLongestSubstring("ab"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// 这种写法不错，但执行次数多。
func lengthOfLongestSubstring(s string) (int, string) {
	m := [128]int{}
	left, right, count := 0, 0, 0
	longest := 0
	longestLeft, longestRight := 0, 0
	for right < len(s) {
		count++
		if m[s[right]] == 0 {
			m[s[right]]++
			right++
		} else {
			m[s[left]]--
			left++
		}
		length := right - left
		if length > longest {
			longest = length
			longestLeft = left
			longestRight = right
		}
	}
	fmt.Println(count)
	return longest, s[longestLeft:longestRight]
}
