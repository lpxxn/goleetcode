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

func TestStrLength(t *testing.T) {
	t.Log(longestSubString2("abcabcbb"))
	//t.Log(longestSubString("bbbbb"))
	//t.Log(longestSubString("pwwkew"))
	//fmt.Println(lengthOfLongestSubstring("abcabbzycdebfdb"))
	t.Log(longestSubString("abcabbzycdebfdb"))
	t.Log(longestSubString2("abcabbzycdebfdb"))

}

func longestSubString(s string) int {
	m := [127]int{}
	left := 0
	longest := 0
	for right := 0; right < len(s); right++ {
		idx := m[s[right]]
		if idx > 0 {
			left = max(left, idx)
		}
		m[s[right]] = right + 1
		longest = max(longest, right-left+1)
	}
	return longest
}

func longestSubString2(s string) int {
	m := [127]int{}
	left := 0
	longest, longestLeft, longestRight := 0, 0, 0
	for right := 0; right < len(s); right++ {
		//fmt.Println(string(s[right]))
		idx := m[s[right]]
		if idx > 0 {
			left = max(left, idx)
		}
		m[s[right]] = right + 1
		currentLongest := max(longest, right-left+1)
		if currentLongest > longest {
			longest = currentLongest
			longestLeft = left
			longestRight = right + 1
		}
	}
	fmt.Println(s[longestLeft:longestRight])
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
