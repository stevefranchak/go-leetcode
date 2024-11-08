package add_two_numbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return prettyPrint(l, 0)
}

func prettyPrint(l *ListNode, depth int) string {
	var prefix string
	if depth > 0 {
		prefix = " -> "
	}

	if l == nil {
		return fmt.Sprintf("%snil", prefix)
	}

	return fmt.Sprintf("%s%d%s", prefix, l.Val, prettyPrint(l.Next, depth+1))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var prev *ListNode
	var carry int

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}

		curr := new(ListNode)
		if head == nil {
			head = curr
		}
		curr.Val = sum
		if prev != nil {
			prev.Next = curr
		}
		prev = curr
	}

	return head
}
