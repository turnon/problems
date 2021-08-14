package main

/*
LC#1791 https://leetcode.com/problems/find-center-of-star-graph/
寻找星图的中心
思路：构造邻接表，循环删除只剩一条边的点，最后剩下那个点就是中心。（但题目说了中心与其它点都有边相连，所以其实可以只看第一二条边，找出重复的那个点，就是中心……）
*/
func findCenter(edges [][]int) int {
	str := struct{}{}
	table := make(map[int]map[int]struct{})
	addToTable := func(from, to int) {
		targets, exist := table[from]
		if !exist {
			targets = make(map[int]struct{})
			table[from] = targets
		}
		targets[to] = str
	}

	for _, edge := range edges {
		addToTable(edge[0], edge[1])
		addToTable(edge[1], edge[0])
	}

	toRemove := make([]int, 0, len(table))
	for len(table) > 1 {
		for src, targets := range table {
			if len(targets) == 1 {
				toRemove = append(toRemove, src)
			}
		}
		for _, n := range toRemove {
			for target := range table[n] {
				delete(table[target], n)
			}
			delete(table, n)
		}
		toRemove = toRemove[0:0]
	}

	for n := range table {
		return n
	}
	return 0
}
