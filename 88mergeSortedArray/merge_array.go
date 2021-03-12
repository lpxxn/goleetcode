package mergeSortedArray

func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	last := len(nums1) - 1
	for n >= 0 {
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[last] = nums1[m]
			m--
		} else {
			nums1[last] = nums2[n]
			n--
		}
		last--
	}
}
