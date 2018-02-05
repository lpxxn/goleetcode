package longest_substring

func method2(str string) int {
	runes := []rune(str);
	runeMap := make(map[rune]int)
	longest := 0;
	preLengest := 0;

	for i, v := range runes {
		var length int
		if val, ok := runeMap[v]; !ok || val < i - preLengest {
			length = preLengest + 1;

		} else {
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