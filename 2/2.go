package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers ...
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curr *ListNode = &ListNode{}
	var result = curr
	var carry int = 0

	for l1 != nil && l2 != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = l1.Val + l2.Val + carry

		if curr.Val >= 10 {
			curr.Val = curr.Val - 10
			carry = 1
		} else {
			carry = 0
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		for l2 != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
			curr.Val = l2.Val + carry
			if curr.Val >= 10 {
				curr.Val = curr.Val - 10
				carry = 1
			} else {
				carry = 0
			}
			l2 = l2.Next
		}
	}

	if l2 == nil {
		for l1 != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
			curr.Val = l1.Val + carry
			if curr.Val >= 10 {
				curr.Val = curr.Val - 10
				carry = 1
			} else {
				carry = 0
			}
			l1 = l1.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = carry
		return result.Next
	}

	return result.Next
}

func main() {
	slice1 := []int{1}
	slice2 := []int{9, 9}

	var cl1, l1 *ListNode
	var cl2, l2 *ListNode

	l1 = &ListNode{}
	l2 = &ListNode{}
	cl1 = l1
	cl2 = l2

	for k, i := range slice1 {
		cl1.Val = i
		if k == len(slice1)-1 {
			break
		}
		cl1.Next = &ListNode{}
		cl1 = cl1.Next
	}

	for k, i := range slice2 {
		cl2.Val = i
		if k == len(slice2)-1 {
			break
		}
		cl2.Next = &ListNode{}
		cl2 = cl2.Next
	}

	head1 := l1
	for l1 != nil {
		fmt.Printf("%d ", l1.Val)
		l1 = l1.Next
	}
	fmt.Printf("\n")

	head2 := l2
	for l2 != nil {
		fmt.Printf("%d ", l2.Val)
		l2 = l2.Next
	}
	fmt.Printf("\n")

	l3 := addTwoNumbers(head1, head2)
	for l3 != nil {
		fmt.Printf("%d ", l3.Val)
		l3 = l3.Next
	}
	fmt.Printf("\n")
}
