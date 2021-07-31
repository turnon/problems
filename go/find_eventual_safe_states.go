package main

/*
LC#802 有向图中不会进入到环的点
https://leetcode.com/problems/find-eventual-safe-states
思路：深度优先遍历，并记录当前路径上的点；如果路径上有点与当前点相同，则表示当前点进入了环；如果某个点的下一点会进入环，则该点也会进入环。
*/
func eventualSafeNodes(graph [][]int) []int {
	walkingIdxes := make([]int, len(graph))
	circles := make([]int, len(graph))
	result := make([]int, 0, len(graph))

	var hasCircle func(idx int) bool
	hasCircle = func(idx int) bool {
		if walkingIdxes[idx] == 1 {
			circles[idx] = 1
			return true
		}

		if circles[idx] != 0 {
			return circles[idx] == 1
		}

		walkingIdxes[idx] = 1
		for _, stop := range graph[idx] {
			if hasCircle(stop) {
				circles[idx] = 1
			}
		}
		if circles[idx] == 0 {
			circles[idx] = -1
		}

		walkingIdxes[idx] = 0
		return circles[idx] == 1
	}

	for idx, _ := range graph {
		if !hasCircle(idx) {
			result = append(result, idx)
		}
	}

	return result
}
