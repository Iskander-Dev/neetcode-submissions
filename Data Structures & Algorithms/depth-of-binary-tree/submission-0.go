/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var count int
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	
	count = 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	return count
}
