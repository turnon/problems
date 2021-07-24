package main

/*
LC#797 https://leetcode.com/problems/all-paths-from-source-to-target/
有向无环图中的所有路径
思路：深度优先遍历，并缓存每个点的后续路径以便减少计算。（题目没说清楚图中存在多个终点时怎样处理才算数，坑……）
*/
func allPathsSourceTarget(graph [][]int) [][]int {
	cache := make([][][]int, len(graph))

	end := len(graph) - 1
	endAtMid := 0
	for i, nextStops := range graph {
		if len(nextStops) == 0 {
			endAtMid = i
		}
	}

	var walk func(start int) [][]int
	walk = func(start int) [][]int {
		paths := cache[start]
		if paths != nil {
			return paths
		}

		nextStops := graph[start]
		if len(nextStops) == 0 {
			if start == endAtMid {
				return [][]int{{start}}
			} else {
				return [][]int{}
			}
		}

		paths = [][]int{}
		for _, next := range nextStops {
			pathsFromNext := walk(next)
			for _, pathFromNext := range pathsFromNext {
				pathFromStart := make([]int, 0, len(pathFromNext)+1)
				pathFromStart = append(pathFromStart, start)
				pathFromStart = append(pathFromStart, pathFromNext...)
				paths = append(paths, pathFromStart)
			}
		}

		cache[start] = paths
		return paths
	}

	result := walk(0)

	if end != endAtMid {
		for i, path := range result {
			if path[len(path)-1] == endAtMid {
				result[i] = path[:len(path)-1]
			}
		}
	}

	return result
}
