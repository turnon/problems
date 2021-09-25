package main

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var maxDepthOfChildren int
	for _, child := range root.Neighbors {
		mxd := maxDepth(child)
		if mxd > maxDepthOfChildren {
			maxDepthOfChildren = mxd
		}
	}
	return 1 + maxDepthOfChildren
}
