package test

import "testing"

/*
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

4656 + 5345 =
*/

func TestAddTwoNum(t *testing.T) {
	a, b := newNum(342), newNum(465)
	c := addTowNun(a, b)
	t.Log(c.toInt())

	a, b = newNum(4656), newNum(5345)
	c = addTowNun(a, b)
	t.Log(c.toInt())
}

type Num struct {
	val  int
	next *Num
}

func newNum(v int) *Num {
	if v == 0 {
		return nil
	}
	current := &Num{val: v % 10}
	current.next = newNum(v / 10)
	return current
}

func (n *Num) toInt() int {
	if n == nil {
		return 0
	}
	return n.val + n.next.toInt()*10
}

func addTowNun(a, b *Num) *Num {
	carry := 0
	rev := &Num{}
	current := rev
	for a != nil || b != nil {
		v1, v2 := 0, 0
		if a != nil {
			v1 = a.val
			a = a.next
		}
		if b != nil {
			v2 = b.val
			b = b.next
		}
		sum := v1 + v2 + carry
		current.next = &Num{val: sum % 10}
		carry = sum / 10
		current = current.next
	}
	if carry > 0 {
		current.next = &Num{val: carry}
	}
	return rev.next
}

/*  24
type NumList struct {
	val  int
	next *NumList
}

func newNumList(val int) *NumList {
	if val == 0 {
		return nil
	}
	rev := &NumList{val: val % 10}
	rev.next = newNumList(val / 10)
	return rev
}

func (n *NumList) ToInt() int {
	if n == nil {
		return 0
	}
	return n.val + n.next.ToInt()*10
}

func addTwoNum(a, b *NumList) *NumList {
	carry := 0
	rev := &NumList{}
	current := rev
	for a != nil || b != nil {
		vala := 0
		valb := 0
		if a != nil {
			vala = a.val
			a = a.next
		}
		if b != nil {
			valb = b.val
			b = b.next
		}
		sum := vala + valb + carry
		carry = sum / 10
		current.next = &NumList{val: sum % 10}
		current = current.next
	}
	if carry > 0 {
		current.next = &NumList{val: carry}
	}
	return rev.next
}

func TestAddTwoNum(t *testing.T) {
	a := newNumList(342)
	b := newNumList(465)
	c := addTwoNum(a, b)
	t.Log(c.ToInt())
	a, b = newNumList(5345), newNumList(4656)
	c = addTwoNum(a, b)
	t.Log(c.ToInt())
}


*/
//type Num struct {
//	val  int
//	next *Num
//}
//
//func NewNum(val int) *Num {
//	if val == 0 {
//		return nil
//	}
//	rev := &Num{}
//	next := rev
//	for val > 0 {
//		next.val = val % 10
//		val = val / 10
//		if val > 0 {
//			next.next = &Num{}
//			next = next.next
//		}
//	}
//	return rev
//}
//
//func (n *Num) ToInt() int {
//	if n == nil {
//		return 0
//	}
//	val := n.val
//	for n.next != nil {
//		val = val*10 + n.next.val
//		n = n.next
//	}
//	return val
//}
//
//func AddTowNum(a, b *Num) *Num {
//	return addTwoNum(a, b, 0)
//}
//
//func addTwoNum(a, b *Num, carry int) *Num {
//	if a == nil && b == nil {
//		if carry > 0 {
//			return &Num{val: carry}
//		}
//		return nil
//	}
//	aValue, bValue := 0, 0
//	if a != nil {
//		aValue = a.val
//	}
//	if b != nil {
//		bValue = b.val
//	}
//	next := &Num{}
//	sum := aValue + bValue + carry
//	next.val = sum % 10
//	next.next = addTwoNum(a.next, b.next, sum/10)
//	return next
//}
//
//func TestLinkFunc(t *testing.T) {
//	a, b := NewNum(342), NewNum(465)
//	c := AddTowNum(a, b)
//	t.Log(c.ToInt())
//
//	a, b = NewNum(5345), NewNum(4656)
//	c = AddTowNum(a, b)
//	t.Log(c.ToInt())
//}
