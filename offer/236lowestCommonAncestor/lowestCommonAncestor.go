package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	var res *TreeNode
	if left != nil && right != nil {
		res = root
	}
	if left != nil && right == nil {
		res = left
	}
	if left == nil && right != nil {
		res = right
	}
	return res
}
