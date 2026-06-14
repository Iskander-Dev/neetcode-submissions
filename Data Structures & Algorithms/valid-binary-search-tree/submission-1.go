/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(root *TreeNode, minValue, maxValue int) bool {
	if root == nil {
		return true
	}
	if root.Val <= minValue || root.Val >= maxValue {
		return false
	}
	return isValid(root.Left, minValue, root.Val) && isValid(root.Right, root.Val, maxValue)
}
