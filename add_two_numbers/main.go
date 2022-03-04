package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

func main() {
	tcs := []struct {
		l1_values  []int
		l2_values  []int
		out_values []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, tc := range tcs {
		l1Nodes := createNodes(tc.l1_values)
		l2Nodes := createNodes(tc.l2_values)

		o := addTwoNumbers(l1Nodes, l2Nodes)
		fmt.Printf("%+v", o)
	}

}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func createNodes(values []int) *ListNode {
	var node *ListNode

	var currentNode *ListNode
	for _, value := range values {
		if node == nil {
			node = &ListNode{
				Val: value,
			}
			currentNode = node
		} else {
			currentNode.Next = &ListNode{
				Val: value,
			}
			currentNode = currentNode.Next
		}
	}

	return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nextNode1 *ListNode
	var nextNode2 *ListNode
	if l1 != nil {
		nextNode1 = l1.Next
	}
	if l2 != nil {
		nextNode2 = l2.Next
	}

	var returnNode *ListNode

	node := &ListNode{
		Val:  0,
		Next: returnNode,
	}

	if l1 != nil {
		node.Val = node.Val + l1.Val
	}

	if l2 != nil {
		node.Val = node.Val + l2.Val
	}

	reminder := 0
	if node.Val > 9 {
		node.Val = node.Val % 10
		reminder = 1
	}

	if nextNode1 != nil {
		nextNode1.Val = nextNode1.Val + reminder
		reminder = 0
	}

	if nextNode2 != nil {
		nextNode2.Val = nextNode2.Val + reminder
		reminder = 0
	}

	if nextNode1 != nil || nextNode2 != nil {
		node.Next = addTwoNumbers(nextNode1, nextNode2)
	} else if reminder > 0 {
		node.Next = &ListNode{
			Val:  reminder,
			Next: nil,
		}
	}

	return node
}
