package test

import "testing"

/*
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

type Num struct {
	val  int
	next *Num
}

func NewNum(val int) *Num {
	if val == 0 {
		return nil
	}
	rev := &Num{}
	next := rev
	for val > 0 {
		next.val = val % 10
		val = val / 10
		if val > 0 {
			next.next = &Num{}
			next = next.next
		}
	}
	return rev
}

func (n *Num) ToInt() int {
	if n == nil {
		return 0
	}
	val := n.val
	for n.next != nil {
		val = val*10 + n.next.val
		n = n.next
	}
	return val
}

func AddTowNum(a, b *Num) *Num {
	return addTwoNum(a, b, 0)
}

func addTwoNum(a, b *Num, carry int) *Num {
	if a == nil && b == nil {
		if carry > 0 {
			return &Num{val: carry}
		}
		return nil
	}
	aValue, bValue := 0, 0
	if a != nil {
		aValue = a.val
	}
	if b != nil {
		bValue = b.val
	}
	next := &Num{}
	sum := aValue + bValue + carry
	next.val = sum % 10
	next.next = addTwoNum(a.next, b.next, sum/10)
	return next
}

func TestLinkFunc(t *testing.T) {
	a, b := NewNum(342), NewNum(465)
	c := AddTowNum(a, b)
	t.Log(c.ToInt())

	a, b = NewNum(5345), NewNum(4656)
	c = AddTowNum(a, b)
	t.Log(c.ToInt())
}
