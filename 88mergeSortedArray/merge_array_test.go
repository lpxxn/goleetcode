package mergeSortedArray

import (
	"reflect"
	"testing"
)

func TestMergeArray(t *testing.T) {
	a1 := []int{1, 2, 3, 0, 0, 0}
	a2 := []int{0, 5, 6}
	MergeTwoArray(a1, 3, a2, 3)
	t.Log(a1)
	a3 := []int{1, 2, 3, 0, 0, 0}
	mergeArr(a3, 3, a2, 3)
	t.Log(reflect.DeepEqual(a1, a3))

}

func mergeArr(a []int, m int, b []int, n int) {
	m--
	n--
	lenA := len(a) - 1
	for n >= 0 {
		if m >= 0 && a[m] > b[n] {
			a[lenA] = a[m]
			//lenA--
			m--
		} else {
			//if a[m] < b[n] {
			a[lenA] = b[n]
			//lenA--
			n--
		}
		lenA--
	}
}

/*
下面是javascript的
双指针 - 从后向前
思路: 开局先让m和n各自减1, 因为数组索引是从0开始的, 然后索引i用来从后向前遍历nums1

! 第一次循环
     m=2
[1,2,3,0,0,0] i=5
[0,5,6]
     n=2
! 因为nums1[2]<nums2[2], 所以nums1[i] = nums2[n], i--, n--

! 第二次循环
     m=2
[1,2,3,0,0,6] i=4
[0,5,6]
   n=1
! 此时nums1[2]<nums2[1], 所以nums1[i] = nums2[n], i--, n--

! 第三次循环
     m=2
[1,2,3,0,5,6] i=3
[0,5,6]
 n=0
! 此时nums1[2]>nums2[0], 所以nums1[i] = nums1[m], i--, m--

! 第四次循环
   m=1
[1,2,3,3,5,6] i=2
[0,5,6]
 n=0
! 此时nums1[1]>nums2[0], 所以nums1[i] = nums1[m], i--, m--

! 第五次循环
 m=0
[1,2,2,3,5,6] i=1
[0,5,6]
 n=0
! 此时nums1[0]>nums2[0], 所以nums1[i] = nums1[m], i--, m--

! 第六次循环
m=-1
[1,1,2,3,5,6] i=0
[0,5,6]
 n=0
! 此时nums1[m]和nums2[n]进行比较
- 到了这里最后一次循环的时候, 要么m=-1, 要么n=-1
- 那我们可以让n>=0来成为while循环的条件
- 当n=-1的说明nums1[0]的位置是正确的, 此时也就无需进行调整, 结束循环即可
- 当m = -1的时候, nums2[-1]会得到一个undefined, undefined跟数字进行比较的时候, 会转换成NaN
- NaN (Not a Number) 跟任何数字进行比较的时候, 都会得到一个 false, 也就意味着必定进入else
- 那我们正好可以把nums2[0]的值赋给nums1[0], 然后正好结束循环

! 推理结束, n 变成 -1 结束循环
[0,1,2,3,5,6]


const merge = (nums1, m, nums2, n) => {
  let i = nums1.length - 1
  m--
  n--
  while (n >= 0) {
    if (nums1[m] > nums2[n]) {
      nums1[i--] = nums1[m--]
    } else {
      nums1[i--] = nums2[n--]
    }
  }
}
*/
