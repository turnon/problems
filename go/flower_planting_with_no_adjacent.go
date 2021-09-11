package main

/*
LC#1042 相邻的花园种不同的花
https://leetcode.com/problems/flower-planting-with-no-adjacent/
思路：先构造双向的邻接表，然后递归地从第1个花园种到第N个，过程中记录每个花园种下的花，并通过邻接表检查准备种的花与旁边的花是否相同
*/
func gardenNoAdj(n int, paths [][]int) []int {
	planting := make([]int, n+1)
	flowers := [4]int{1, 2, 3, 4}

	gardenAdjs := make([][]int, n+1)
	for _, path := range paths {
		gardenAdjs[path[0]] = append(gardenAdjs[path[0]], path[1])
		gardenAdjs[path[1]] = append(gardenAdjs[path[1]], path[0])
	}

	canPlant := func(g int, f int) bool {
		for _, adj := range gardenAdjs[g] {
			if planting[adj] == f {
				return false
			}
		}
		return true
	}

	var doPlant func(g int) bool
	doPlant = func(g int) bool {
		if g > n {
			return true
		}
		if planting[g] != 0 {
			return doPlant(g + 1)
		}
		for _, f := range flowers {
			if canPlant(g, f) {
				planting[g] = f
				if doPlant(g + 1) {
					return true
				}
			}
		}
		planting[g] = 0
		return false
	}

	doPlant(1)
	return planting[1:]
}
