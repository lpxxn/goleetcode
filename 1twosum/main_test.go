package main

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 5, 18, 21}
	revArr := twoSum(nums, 23)
	t.Log(revArr)

	revArr2 := myTwoSum([]int64{1, 2, 5, 18, 21}, 23)
	t.Log(revArr2)
}

func myTwoSum(a []int64, s int64) []int {
	m := map[int64]int{}
	for i, v := range a {
		if pIdx, ok := m[s-v]; ok {
			return []int{pIdx, i}
		}
		m[v] = i
	}
	return nil
}
