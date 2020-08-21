package main

import "sort"

func main() {

}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsLen1, numsLen2 := len(nums1), len(nums2)
	numsTotal := numsLen1 + numsLen2
	needMiddleTwo := false

	if numsTotal % 2 == 0 {
		needMiddleTwo = true
	}

	newArr := append(nums1, nums2...)
	sort.Ints(newArr)

	index := numsTotal/2
	middle1 := newArr[index]
	middle2 := 0
	if needMiddleTwo {
		middle2 = newArr[index-1]
		return float64(middle1 + middle2) / 2
	}

	return float64(middle1)
}
