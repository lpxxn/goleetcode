package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 3, 4, 5, 6}
	median := findMedianSortedArrays2(arr1, arr2)
	println(median)

	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4, 5}))
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	var nums []int
	for i < len(nums1) || j < len(nums2) {
		if i >= len(nums1) {
			nums = append(nums, nums2[j])
			j++
		} else if j >= len(nums2) {
			nums = append(nums, nums1[i])
			i++
		} else if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		return float64(nums[totalLen/2]+nums[totalLen/2-1]) / 2
	} else {
		return float64(nums[totalLen/2])
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsLen1, numsLen2 := len(nums1), len(nums2)
	numsTotal := numsLen1 + numsLen2
	needMiddleTwo := false

	if numsTotal%2 == 0 {
		needMiddleTwo = true
	}

	newArr := append(nums1, nums2...)
	sort.Ints(newArr)

	index := numsTotal / 2
	middle1 := newArr[index]
	middle2 := 0
	if needMiddleTwo {
		middle2 = newArr[index-1]
		return float64(middle1+middle2) / 2
	}

	return float64(middle1)
}
