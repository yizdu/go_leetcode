package leetcode

func dfs104(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}(dfs104(root.Left), dfs104(root.Right)) + 1
}

func MaxDepth0(root *TreeNode) int {

	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		levelWidth := len(queue)

		for i := 0; i < levelWidth; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth += 1
	}
	return depth

}
