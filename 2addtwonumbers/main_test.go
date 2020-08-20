package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	v := IToNum(342)
	fmt.Println(v, v.ToInt())
	v = AddTwoNum(IToNum(342), IToNum(465))
	fmt.Println(v, v.ToInt())

	v = AddTwoNum(IToNum(42), IToNum(65))
	fmt.Println(v, v.ToInt())
}

type Num struct {
	Val  int
	Next *Num
}

func (n Num) String() string {
	if n.Next == nil {
		return fmt.Sprintf("%d", n.Val)
	}
	return fmt.Sprintf("%d -> %s", n.Val, n.Next.String())
}

func (n Num) IntStr() string {
	if n.Next == nil {
		return fmt.Sprintf("%d", n.Val)
	}
	return fmt.Sprintf("%s%d", n.Next.IntStr(), n.Val)
}

func (n Num) ToInt() int {
	str := n.IntStr()
	v, _ := strconv.Atoi(str)
	return v
}

func IToNum(v int) *Num {
	if v == 0 {
		return &Num{Val: 0}
	}
	newNum := &Num{}
	nextNum := newNum
	for v > 0 {
		nextNum.Next = &Num{Val: v % 10}
		nextNum = nextNum.Next
		v = v / 10
	}
	return newNum.Next
}

func AddTwoNum(a, b *Num) *Num {
	newNum := &Num{}
	var currentNum *Num
	currentNum = newNum
	carry := 0
	for a != nil || b != nil {
		x, y := 0, 0
		if a != nil {
			x = a.Val
			a = a.Next
		}
		if b != nil {
			y = b.Val
			b = b.Next
		}
		sum := x + y + carry
		currentNum.Next = &Num{Val: sum % 10}
		carry = sum / 10
		currentNum = currentNum.Next
	}
	if carry > 0 {
		currentNum.Next = &Num{Val: carry}
	}
	return newNum.Next
}
