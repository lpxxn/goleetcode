package longest_substring

import (
	"fmt"
	"strings"
)

func longestSubString(s string) int {
	m := [128]int{}
	longest := 0
	left := 0
	for right := 0; right < len(s); right++ {
		// s[right] 是新遇到的字母
		// m 是记录字母对应的新位置，
		// 所以 m[s[right]] 就是新字母对应的新位置
		left = max(left, m[s[right]])
		m[s[right]] = right + 1
		longest = max(longest, right-left+1)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	var strs string = ""
	var max int = 0
	for _, item := range s {
		lenStrs := len(strs)
		value := string(item)
		if lenStrs > 0 && strings.Contains(strs, value) {
			if max < lenStrs {
				max = lenStrs
			}
			index := strings.Index(strs, value)
			strs = strs[index+1:]
		}
		strs += value
	}
	lens := len(strs)
	if lens > max {
		return lens
	}
	return max
}

// 这个感觉不好
func method2(str string) int {
	runes := []rune(str)
	runeMap := make(map[rune]int)
	longest := 0
	preLongest := 0

	for i, v := range runes {
		var length int
		currentV := string(v)
		// ok = true i - preLongest 是上一个重复的位置abcabbzycdebfdb  每二个c 8 - 3 = 5 也就是b的位置，然后每一个c的位置是 2
		// 也就是说，最近的一个重复的位置比 c 要大。
		// 也可以这么说，就是两个相同的字母 每一个c(2)和第二个c(8)之前有没有新的多个重复的字母（a重复 bca b重复 cab b重复 b）
		if val, ok := runeMap[v]; !ok || val < i-preLongest {
			fmt.Println(currentV, val)
			length = preLongest + 1
		} else {
			length = i - val
		}
		if length > longest {
			longest = length
		}
		preLongest = length
		runeMap[v] = i
	}
	return longest
}
