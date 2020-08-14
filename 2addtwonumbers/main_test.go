package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	v := IToNum(342)
	fmt.Println(v)
	fmt.Println(v.ToInt())
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
	return nil
}
