// Problem: https://leetcode.com/problems/path-sum-iii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countPathSumEqual(curPath *[]int, targetSum int) int {
	sum := 0

	count := 0

	for i := len(*curPath) - 1; i >= 0; i-- {
		sum += (*curPath)[i]

		if sum == targetSum {
			count++
		}
	}

	return count
}

func helper(root *TreeNode, curPath *[]int, targetSum int) int {
	if root == nil {
		return 0
	}

	*curPath = append(*curPath, root.Val)

	ans := countPathSumEqual(curPath, targetSum)

	ans += helper(root.Left, curPath, targetSum)

	ans += helper(root.Right, curPath, targetSum)

	*curPath = (*curPath)[:len(*curPath)-1]

	return ans
}

func pathSum(root *TreeNode, targetSum int) int {
	return helper(root, &[]int{}, targetSum)
}