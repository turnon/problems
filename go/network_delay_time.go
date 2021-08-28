package main

/*
LC#743 网络延时
https://leetcode.com/problems/network-delay-time
思路：构造邻接表。然后从给定结点开始遍历，并记录（或者递归更新）每个结点的最短用时。最后找出用时最长的结点，即整体用时。如果某个结点没有记录用时，则没有完全遍历。
*/
func networkDelayTime(times [][]int, n int, k int) int {
	type node struct {
		id    int
		delay int
	}

	type costTime struct {
		time int
	}

	// 构造邻接表
	adjacent := make([][]*node, n+1)

	for _, time := range times {
		from, to, delay := time[0], time[1], time[2]
		adjacent[from] = append(adjacent[from], &node{id: to, delay: delay})
	}

	// 从给定结点开始遍历，并记录（或者递归更新）每个结点的最短用时
	costTimes := make([]*costTime, n+1)

	var walk func(nodeId int, costing int)
	walk = func(nodeId int, costed int) {
		cost := costTimes[nodeId]
		if cost == nil {
			cost = &costTime{time: 100}
			costTimes[nodeId] = cost
		}

		if cost.time > costed {
			cost.time = costed
			for _, adjacentNode := range adjacent[nodeId] {
				walk(adjacentNode.id, adjacentNode.delay+costed)
			}
		}
	}

	walk(k, 0)

	// 找出用时最长的结点，即整体用时。如果某个结点没有记录用时，则没有完全遍历。
	maxCost := 0
	for _, cost := range costTimes[1:] {
		if cost != nil {
			if cost.time > maxCost {
				maxCost = cost.time
			}
		} else {
			return -1
		}
	}

	return maxCost
}
