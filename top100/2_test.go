package leetcode_top100

import (
	"strconv"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 初始化一个新的链表
	head := &ListNode{}
	// 2. 初始化一个指针，始终指向链表的最后一个节点
	tail := head

	for l1 != nil || l2 != nil {
		l1Val, l2Val := 0, 0

		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		divisor, remainder := (l1Val+l2Val+tail.Val)/10, (l1Val+l2Val+tail.Val)%10

		tail.Val = remainder
		if l1 != nil || l2 != nil || divisor != 0 {
			tail.Next = &ListNode{Val: divisor}
			tail = tail.Next
		}
	}

	return head
}

func Test_AddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	result := addTwoNumbers(l1, l2)
	resultStr := ""
	for result != nil {
		resultStr += strconv.Itoa(result.Val)
		t.Log(result.Val)
		result = result.Next
	}
	if resultStr != "89991" {
		t.Errorf("Expected: 89991, got: %s", resultStr)
	}
}
