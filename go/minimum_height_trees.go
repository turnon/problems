package main

/*
LC#310 https://leetcode.com/problems/minimum-height-trees
图中高度最低的树
思路：将边数组转换成邻接表，然后计算每个点的深度，并在过程中缓存每个方向的最大长度，最后筛选出深度最低的点。
论坛的做法是在循环中不断排除只剩一条边的点，最终剩余的就是核心部分距离外围很均匀的那些点，想法不一般……
*/
func findMinHeightTrees(n int, edges [][]int) []int {
	edgesMap := make([][]int, n)
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		edgesMap[n1] = append(edgesMap[n1], n2)
		edgesMap[n2] = append(edgesMap[n2], n1)
	}

	already := &struct{}{}
	walkedDepths := make(map[[2]int]int)

	var walk func(node int, walked []*struct{}) int
	walk = func(node int, walked []*struct{}) int {
		walked[node] = already
		depth := 0
		for _, next := range edgesMap[node] {
			if walked[next] == already {
				continue
			}
			direction := [2]int{node, next}
			newDepth := walkedDepths[direction]
			if newDepth == 0 {
				newDepth = walk(next, walked) + 1
				walkedDepths[direction] = newDepth
			} else {
				walked[next] = already
			}
			if newDepth > depth {
				depth = newDepth
			}
		}
		walked[node] = nil
		return depth
	}

	depths := make([]int, n)
	for node, _ := range edgesMap {
		depths[node] = walk(node, make([]*struct{}, n))
	}

	nodes := []int{}
	mindepth := n
	for node, depth := range depths {
		if depth < mindepth {
			nodes = nodes[0:0]
			mindepth = depth
		}
		if depth == mindepth {
			nodes = append(nodes, node)
		}
	}

	return nodes
}
