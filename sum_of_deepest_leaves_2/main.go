package main

import "fmt"

/*

Given the root of a binary tree, return the sum of values of its deepest leaves.

Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15

Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 19

*/

func main() {

	root := &TreeNode{
		Val: 6,
	}
	root.Left = &TreeNode{
		Val: 7,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Left.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Right = &TreeNode{
		Val: 7,
	}
	root.Right.Left = &TreeNode{
		Val: 1,
	}
	root.Right.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Left.Left = &TreeNode{
		Val: 9,
	}
	root.Left.Right.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Right.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Right.Right = &TreeNode{
		Val: 5,
	}
	root.Right.Right.Right.Left = &TreeNode{
		Val: 5,
	}
	root.Right.Right.Right.Right = &TreeNode{
		Val: 5,
	}

	fmt.Println(deepestLeavesSum(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(root *TreeNode, level int, maxLevel map[int]int, res map[int]*TreeNode) {

	if root != nil {
		level++
		find(root.Left, level, maxLevel, res)

		if level > maxLevel[0] {
			maxLevel[0] = level

			for k := range res {
				delete(res, k)
			}

			res[0] = root
		} else if level == maxLevel[0] {
			res[len(res)] = root
		}
		find(root.Right, level, maxLevel, res)
	}
}

func deepestLeavesSum(root *TreeNode) int {
	res := make(map[int]*TreeNode)
	var maxLevel = map[int]int{
		0: 0,
	}
	var level int = 0

	find(root, level, maxLevel, res)

	var amount = 0

	for _, r := range res {
		amount = amount + r.Val
	}
	return amount
}
