package longest_substring

import "fmt"

func method2(str string) int {
	runes := []rune(str);
	runeMap := make(map[rune]int)
	longest := 0;
	preLengest := 0;

	for i, v := range runes {
		var length int
		if val, ok := runeMap[v]; !ok || val < i - preLengest {
			fmt.Println("val :", val, " i: ", i, )
			length = preLengest + 1;

		} else {
			fmt.Println("val :", val, " i: ", i)
			length = i - val
		}

		if length > longest {
			longest = length
		}

		preLengest = length
		runeMap[v] = i
	}
	return longest
}