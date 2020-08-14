package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	v := IToNum(342)
	fmt.Println(v)
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

func IToNum(v int) *Num {
	if v == 0 {
		return &Num{Val: 0}
	}
	newNum := &Num{}
	nextNum := newNum
	for v > 0 {
		nextNum.Next = &Num{Val: v % 10}
		v = v / 10
		nextNum = nextNum.Next
	}
	return newNum.Next
}
