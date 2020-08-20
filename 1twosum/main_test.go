package main

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 5, 18, 21}
	revArr := twoSum(nums, 23)
	t.Log(revArr)

	revArr2 := myTwoSum([]int{1, 2, 5, 18, 21}, 23)
	t.Log(revArr2)
}

func myTwoSum(p []int, total int) []int {
	s := map[int]int{}
	for i, item := range p {
		if idx, ok := s[total-item]; ok {
			return []int{idx, i}
		}
		s[item] = i
	}
	return nil
}
