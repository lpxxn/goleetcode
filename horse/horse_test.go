package horse

import (
	"fmt"
	"testing"
)

func TestHorse1(t *testing.T) {
	horses := []int{5, 8, 2, 9, 3, 1, 7, 6, 4, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64}

	// 分成8组
	for i := 0; i < 8; i++ {
		group := horses[i*8 : (i+1)*8]

		// 找出最快的4匹马
		fastest := make([]int, 4)
		for j := 0; j < 4; j++ {
			fastest[j] = group[j]
		}
		for j := 4; j < len(group); j++ {
			for k := 0; k < 4; k++ {
				if group[j] < fastest[k] {
					fastest[k] = group[j]
					break
				}
			}
		}

		fmt.Println("Fastest horses in group", i+1, ":", fastest)
	}
}
