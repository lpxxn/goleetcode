package longest_substring

import (
	"testing"
	"fmt"
)

func Test_longest(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabbcbb"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("ab"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}


func Test_logest2(t *testing.T) {

	fmt.Println(method2("abcabbcbb"))
	fmt.Println(method2("a"))
	fmt.Println(method2("ab"))
	fmt.Println(method2("dvdf"))
}
