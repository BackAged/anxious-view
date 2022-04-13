// Problem: https://leetcode.com/problems/path-sum-ii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// isLeafNode returns true if the node is a leaf node.
func isLeafNode(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}

func helper(root *TreeNode, targetSum, curSum int, cur *[]int, ans *[][]int) {
	if root == nil {
		return
	}

	*cur = append(*cur, root.Val)

	if curSum+root.Val == targetSum && isLeafNode(root) {
		*ans = append(*ans, append([]int{}, *cur...))
	}

	helper(root.Left, targetSum, curSum+root.Val, cur, ans)
	helper(root.Right, targetSum, curSum+root.Val, cur, ans)

	*cur = (*cur)[:len(*cur)-1]
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	cur := []int{}

	helper(root, targetSum, 0, &cur, &ans)

	return ans
}