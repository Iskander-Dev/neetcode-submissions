/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var valid func(n *TreeNode, min, max int) bool
	valid = func(n *TreeNode, min, max int) bool {
		if n == nil {
			return true
		}
		if !(min < n.Val && n.Val < max) {
			return false
		}
		return (valid(n.Left, min, n.Val) && valid(n.Right, n.Val, max))
	}
	return valid(root, math.MinInt, math.MaxInt)
}
