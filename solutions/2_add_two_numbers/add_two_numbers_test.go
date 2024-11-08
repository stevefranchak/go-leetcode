package add_two_numbers

import (
	"reflect"
	"testing"
)

func makeLinkedListFromSlice(s []int) *ListNode {
	var head *ListNode
	var curr *ListNode
	var prev *ListNode
	for _, n := range s {
		if curr == nil {
			curr = new(ListNode)
		}
		if head == nil {
			head = curr
		}
		curr.Val = n
		if prev != nil {
			prev.Next = curr
		}
		prev = curr
		curr = nil
	}
	return head
}

func TestMakeLinkedListFromSlice(t *testing.T) {
	cases := []struct {
		input []int
		want  *ListNode
	}{
		{[]int{1}, &ListNode{Val: 1}},
		{[]int{1, 2, 3}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
		{[]int{}, nil},
	}

	for _, c := range cases {
		got := makeLinkedListFromSlice(c.input)
		// no more than 100 nodes in a linked list for this problem, not worried about performance
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("makeLinkedListFromSlice(%v) = %s; want %s", c.input, got, c.want)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			makeLinkedListFromSlice([]int{2, 4, 3}),
			makeLinkedListFromSlice([]int{5, 6, 4}),
			makeLinkedListFromSlice([]int{7, 0, 8}),
		},
		{
			makeLinkedListFromSlice([]int{0}),
			makeLinkedListFromSlice([]int{0}),
			makeLinkedListFromSlice([]int{0}),
		},
		{
			makeLinkedListFromSlice([]int{0}),
			makeLinkedListFromSlice([]int{9, 1}),
			makeLinkedListFromSlice([]int{9, 1}),
		},
		{
			makeLinkedListFromSlice([]int{9, 9, 9, 9, 9, 9, 9}),
			makeLinkedListFromSlice([]int{9, 9, 9, 9}),
			makeLinkedListFromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, c := range cases {
		got := addTwoNumbers(c.l1, c.l2)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("addTwoNumbers(%s, %s) = %s; want %s", c.l1, c.l2, got, c.want)
		}
	}
}
