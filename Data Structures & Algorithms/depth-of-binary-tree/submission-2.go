/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth, rightDepth := maxDepth(root.Left), maxDepth(root.Right)
    return max(leftDepth, rightDepth) + 1
}
