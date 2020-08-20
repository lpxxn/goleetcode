package longest_substring

import (
	"fmt"
	"testing"
)

func Test_longest(t *testing.T) {
	fmt.Println(longestSubString("abcabbzycdebfdb"))
	fmt.Println(longestSubString("a"))
	fmt.Println(longestSubString("aaaa"))
	fmt.Println(longestSubString("ab"))
	fmt.Println(longestSubString("dvdf"))
	fmt.Println(longestSubString("pwwkew"))
}

func Test_longestWithStr(t *testing.T) {
	fmt.Println(longestSubString2("abcabbzycdebfdb"))
	fmt.Println(longestSubString2("abcabcbb"))
	fmt.Println(longestSubString2("a"))
	fmt.Println(longestSubString2("aaaa"))
	fmt.Println(longestSubString2("ab"))
	fmt.Println(longestSubString2("dvdf"))
	fmt.Println(longestSubString2("pwwkew"))
}

func Test_longest1(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabbcbb"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("ab"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func Test_longest2(t *testing.T) {
	fmt.Println(method2("abcabbzycdebfdb"))
	fmt.Println(method2("abcaefghcdebfdb"))
	fmt.Println(method2("a"))
	fmt.Println(method2("ab"))
	fmt.Println(method2("dvdf"))
}


func myLongest(s string) (int, string) {
	a := [128]int{}
	longest := 0
	maxLeftIdx, maxRightIdx := 0, 0
	for i, v := range s {

	}
}