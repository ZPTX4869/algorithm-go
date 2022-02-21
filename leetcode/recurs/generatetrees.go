package recurs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return make([]*TreeNode, 0)
	}

	roots := generate(1, n+1)

	return roots
}

func generate(s, e int) []*TreeNode {
	if s == e {
		return []*TreeNode{nil}
	}

	roots := make([]*TreeNode, 0)
	for i := s; i < e; i++ {
		leftNodes := generate(s, i)
		rightNodes := generate(i+1, e)

		for _, left := range leftNodes {
			for _, right := range rightNodes {
				root := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}

				roots = append(roots, root)
			}
		}
	}

	return roots
}
