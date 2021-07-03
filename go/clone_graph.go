package main

/*
LC#133 https://leetcode.com/problems/clone-graph
复制图
思路：图深度优先遍历，并使用一个map（或根据题目条件创建一个长度101的数组）记录已经过的点。
*/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldNewMap := make([]*Node, 101)

	var _cloneGraph func(n *Node) *Node
	_cloneGraph = func(n *Node) *Node {
		newNode := oldNewMap[n.Val]
		if newNode != nil {
			return newNode
		}

		newNode = &Node{Val: n.Val}
		oldNewMap[n.Val] = newNode

		newNode.Neighbors = make([]*Node, 0, len(n.Neighbors))
		for _, oldNeighour := range n.Neighbors {
			newNeighour := _cloneGraph(oldNeighour)
			newNode.Neighbors = append(newNode.Neighbors, newNeighour)
		}
		return newNode
	}
	return _cloneGraph(node)
}

func _cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldNewMap := make(map[*Node]*Node)

	var _cloneGraph func(n *Node) *Node
	_cloneGraph = func(n *Node) *Node {
		newNode, exists := oldNewMap[n]
		if exists {
			return newNode
		}

		newNode = &Node{Val: n.Val}
		oldNewMap[n] = newNode

		for _, oldNeighour := range n.Neighbors {
			newNeighour := _cloneGraph(oldNeighour)
			newNode.Neighbors = append(newNode.Neighbors, newNeighour)
		}
		return newNode
	}
	return _cloneGraph(node)
}
