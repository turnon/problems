package main

/*
LC#684 https://leetcode.com/problems/redundant-connection
多余的连线
思路：构造邻接表，每读取一条连线，就检查它的两个端点是否已经联通（递归检查两点在邻接表中是否相邻或者与邻接点相邻）
如已联通则返回该连线，否则将两点加入邻接表。
论坛使用并查集算法，能减少很多跳转。
*/
func findRedundantConnection(edges [][]int) []int {
	adjecent := make([][]int, len(edges)+1)

	var isLinked func(n1, n2 int, visited map[int]bool) bool
	isLinked = func(n1, n2 int, visited map[int]bool) bool {
		if visited[n1] && visited[n2] {
			return false
		} else {
			visited[n1], visited[n2] = true, true
		}
		for _, linkingNode := range adjecent[n1] {
			if linkingNode == n2 || isLinked(linkingNode, n2, visited) {
				return true
			}
		}
		return false
	}

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		if isLinked(n1, n2, make(map[int]bool)) {
			return edge
		}
		adjecent[n1] = append(adjecent[n1], n2)
		adjecent[n2] = append(adjecent[n2], n1)
	}

	return nil
}
