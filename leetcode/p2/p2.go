/*
	2. Add Two Numbers

	You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

	You may assume the two numbers do not contain any leading zero, except the number 0 itself.

	Example

	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.

	**********************************************************

	2. 两数相加

	给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

	你可以假设除了数字 0 之外，这两个数字都不会以零开头。

	示例：

	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
*/

package p2

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	初等数学
	时间复杂度O(max(m,n))
	空间复杂度O(max(m,n)):max(m,n)+1
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lr := &ListNode{}
	lrc := lr
	carry := 0
	for {
		sum := l1.Val + l2.Val + carry
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}
		lrc.Val = sum

		if l1.Next == nil && l2.Next == nil && carry == 0 {
			break
		}
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{0, nil}
		}
		l1 = l1.Next
		l2 = l2.Next
		lrc.Next = &ListNode{0, nil}
		lrc = lrc.Next
	}
	return lr
}

/*
	最快
	初等数学
	时间复杂度O(max(m,n))
	空间复杂度O(max(m,n)):max(m,n)+1

	累加代替情况判断
*/
func AddTwoNumbersFastest(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	beforeHeadNode := ListNode{}
	lastNode := &beforeHeadNode
	for l1 != nil || l2 != nil || carry != 0 {
		a := carry
		carry = 0
		if l1 != nil {
			a += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			a += l2.Val
			l2 = l2.Next
		}
		if a > 9 {
			a = a - 10
			carry = 1
		}
		currentNode := ListNode{
			Val:  a,
			Next: nil,
		}
		lastNode.Next = &currentNode
		lastNode = &currentNode
	}
	return beforeHeadNode.Next
}
