package test

import (
	"fmt"
	"testing"
)

/*
例子： 342 + 465 = 807.
转换: (2 -> 4 -> 3) + (5 -> 6 -> 4)
每一位相加，过10进1
结果: 7 -> 0 -> 8 = 807
eg:
4656 + 5345 = 10001
*/
type numValue struct {
	value int
	next  *numValue
}

func NewNumValue(value int) *numValue {
	if value == 0 {
		return nil
	}
	rev := &numValue{value: value % 10}
	rev.next = NewNumValue(value / 10)
	return rev
}
func (n *numValue) String() string {
	if n.next == nil {
		return fmt.Sprintf("%d", n.value)
	}
	return fmt.Sprintf("%d -> %s", n.value, n.next.String())
}

func (n *numValue) ToInt() int {
	if n == nil {
		return 0
	}
	return n.next.ToInt()*10 + n.value
}

func AddTowNum(a, b *numValue, carry int) *numValue {
	if a == nil && b == nil {
		return nil
	}
	v1 := 0
	if a != nil {
		v1 = a.value
		a = a.next
	}
	v2 := 0
	if b != nil {
		v2 = b.value
		b = b.next
	}
	sum := v1 + v2 + carry
	carry = sum / 10
	return &numValue{value: sum % 10, next: AddTowNum(a, b, carry)}
}

func TestName(t *testing.T) {
	a := NewNumValue(342)
	t.Log(a.String())
	b := NewNumValue(465)
	t.Log(b.String())

	c := AddTowNum(a, b, 0)
	t.Log(c.String())
	t.Log(c.ToInt())
}
