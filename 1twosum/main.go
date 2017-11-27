package main

import "fmt"

func main() {
	nums := []int {1, 2, 5, 18, 21}
	revArr := twoSum(nums, 20)
	fmt.Print(revArr)
}

// O(N)
func twoSum(nums []int, target int) []int{
	mp := make(map[int]int)

	for i, num := range nums {
		if idx, ok := mp[target - num]; ok{
			return []int{idx, i}
		}
		mp[num] = i
	}
	return  nil
}