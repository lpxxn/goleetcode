package main

import (
	"fmt"
)

// zigzag conversion
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result []string = make([]string, numRows)
	var currentRow int
	var isDown bool
	for _, c := range s {
		result[currentRow] += string(c)
		if currentRow == 0 || currentRow == numRows-1 {
			isDown = !isDown
		}
		if isDown {
			currentRow++
		} else {
			currentRow--
		}
	}
	rev := ""
	for _, item := range result {
		rev += item
	}
	//return strings.Join(result, "")
	return rev
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
