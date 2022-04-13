// Problem: https://leetcode.com/problems/sum-root-to-leaf-numbers/

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

func helper(root *TreeNode, cur int) int {
    if root == nil {
        return 0
    }
    
    if isLeafNode(root) {
        return cur * 10 + root.Val
    }
    
    l := helper(root.Left, cur * 10 + root.Val)
    
    r := helper(root.Right, cur * 10 + root.Val)
    
    return l + r
}

func sumNumbers(root *TreeNode) int {
    return helper(root, 0)
}