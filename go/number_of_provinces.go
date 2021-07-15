package main

/*
LC#547 https://leetcode.com/problems/number-of-provinces
省份数量
思路：每个城市都深度遍历（并在过程中记录经过的城市），除非它已走过。对有作为起点的城市计数。
*/
func findCircleNum(isConnected [][]int) int {
	visited := make([]int, len(isConnected))
	provinces := 0

	var walk func(city int)
	walk = func(city int) {
		for anotherCity, connected := range isConnected[city] {
			if city == anotherCity {
				continue
			}
			if connected == 1 && visited[anotherCity] == 0 {
				visited[anotherCity] = 1
				walk(anotherCity)
			}
		}
	}

	for city, _ := range isConnected {
		if visited[city] == 0 {
			visited[city] = 1
			walk(city)
			provinces += 1
		}
	}

	return provinces
}

func findCircleNum1(isConnected [][]int) int {
	visited := make([]int, len(isConnected))
	provinces := 0

	for city, connecteds := range isConnected {
		if visited[city] == 1 {
			continue
		}
		for anotherCity, connected := range connecteds {
			if city == anotherCity {
				continue
			}
			if connected == 1 {
				visited[anotherCity] = 1
			}
		}
		provinces += 1
	}

	return provinces
}
