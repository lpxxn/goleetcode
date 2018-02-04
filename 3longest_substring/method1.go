package longest_substring

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var strs string = ""
	var max int = 0
	for _, item := range s {
		lenStrs := len(strs)
		value := string(item)
		if lenStrs > 0 && strings.Contains(strs, value) {
			if  max < lenStrs {
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

