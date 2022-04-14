package z_test

import "testing"

func TestName(t *testing.T) {
	t.Log(addTwoNum(newNum(342), newNum(465)).toInt())
	t.Log(addTwoNum(newNum(42), newNum(465)).toInt())
}

type num struct {
	val  int
	next *num
}

func newNum(val int) *num {
	if val == 0 {
		return nil
	}
	current := &num{val: val % 10}
	current.next = newNum(val / 10)
	return current
}

func addTwoNum(v1, v2 *num) *num {
	rev := &num{}
	carry := 0
	current := rev
	for v1 != nil || v2 != nil || carry > 0 {
		a1, a2 := 0, 0
		if v1 != nil {
			a1 = v1.val
			v1 = v1.next
		}

		if v2 != nil {
			a2 = v2.val
			v2 = v2.next
		}
		sum := a1 + a2 + carry
		current.next = &num{val: sum % 10}
		current = current.next
		carry = sum / 10
	}
	return rev.next
}

func (n *num) toInt() int {
	if n == nil {
		return 0
	}
	return n.val + n.next.toInt()*10
}
