package z_test

import "testing"

func TestA(t *testing.T) {
	a := []int{1, 2, 3}
	// remove the first element
	a = a[1:]
	// 从a中删除第一个元素
	a = a[:len(a)-1]
}
