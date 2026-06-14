/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pStack := []*TreeNode{p}
	qStack := []*TreeNode{q}

	for len(pStack) > 0 && len(qStack) > 0 {
		pNode := pStack[len(pStack)-1]
		qNode := qStack[len(qStack)-1]
		pStack = pStack[:len(pStack)-1]
		qStack = qStack[:len(qStack)-1]

		if pNode != nil && qNode != nil {
			if pNode.Val == qNode.Val {
				pStack = append(pStack, pNode.Right)
				pStack = append(pStack, pNode.Left)
				qStack = append(qStack, qNode.Right)
				qStack = append(qStack, qNode.Left)
			} else {
				return false
			}
		} else if pNode == nil && qNode == nil {
			continue
		} else {
			return false
		}
	}

	return true
}
