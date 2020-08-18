package longest_substring

import (
	"testing"
	"fmt"
)


func Test_longest(t *testing.T) {
	fmt.Println(longestSubString("abcabbzycdebfdb"))
	fmt.Println(longestSubString("a"))
	fmt.Println(longestSubString("aaaa"))
	fmt.Println(longestSubString("ab"))
	fmt.Println(longestSubString("dvdf"))
	fmt.Println(longestSubString("pwwkew"))
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
