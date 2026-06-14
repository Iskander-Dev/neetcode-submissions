/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		arr := []int{}
		for range n {
			curr := queue[0]
			queue = queue[1:]
			arr = append(arr, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		res = append(res, arr)
	}
	return res
}
