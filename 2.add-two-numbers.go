/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (29.69%)
 * Total Accepted:    668.6K
 * Total Submissions: 2.3M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Example:
 *
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	p, q := l1, l2
	curr := head
	carry := 0
	var x, y int
	for p != nil || q != nil {
		if p != nil {
			x = p.Val

		} else {
			x = 0

		}
		if q != nil {
			y = q.Val

		} else {
			y = 0

		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = new(ListNode)
		curr.Next.Val = sum % 10
		curr = curr.Next
		if p != nil {
			p = p.Next

		}
		if q != nil {
			q = q.Next

		}

	}
	if carry > 0 {
		curr.Next = new(ListNode)
		curr.Next.Val = carry

	}
	return head.Next
}
