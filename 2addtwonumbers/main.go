package main

import (
	"fmt"
	"errors"
)

func main() {
	l1, err := CreateNode(342)
	if err != nil {
		panic(err)
	}

	//l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	fmt.Println(l1)

	l2, err := CreateNode(465)
	//l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(l2)
	rev := addTwoNumbers(l1, l2)
	fmt.Println(rev)
}


type ListNode struct {
	Val int
	Next *ListNode
}

func (ln ListNode) String() string {
	if ln.Next == nil {
		return fmt.Sprintf("%d", ln.Val)
	} else {
		return fmt.Sprintf("%d -> %s", ln.Val, ln.Next.String())
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	dummyNode := ListNode{0, nil}
	carry := 0
	current := &dummyNode

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}
		sum := x + y + carry
		current.Next = &ListNode{sum % 10, nil}
		carry = sum / 10
		fmt.Println("sum: ", sum, "  sum % 10 = ", sum % 10)
		fmt.Println("carry: ", carry, " is sum/10")
		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		current.Next = &ListNode{1, nil}
	}
	return dummyNode.Next
}

func CreateNode(value int) (*ListNode, error) {
	if value == 0 {
		return &ListNode{0, nil}, nil
	} else if value > 0 {
		node := &ListNode{0, nil}
		currentNode := node
		for value > 0 {
			current := value % 10
			currentNode.Next = &ListNode{current, nil}
			currentNode = currentNode.Next
			value = value / 10
			fmt.Println("value : ", value)
		}
		return node.Next, nil
	} else {
		return nil, errors.New("negative number")
	}
}