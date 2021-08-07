package main

/**
LC#841 解锁全部房间
https://leetcode.com/problems/keys-and-rooms
思路：深度或广度优先遍历，使用数组记录已走过的房间以免爆栈或死循环，最后检查该数组有没有房间未走过。
**/
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]int, len(rooms))

	var visit func(roomId int)
	visit = func(roomId int) {
		if visited[roomId] == 1 {
			return
		}
		visited[roomId] = 1
		for _, nextRonextRoomId := range rooms[roomId] {
			visit(nextRonextRoomId)
		}
	}
	visit(0)

	for _, v := range visited {
		if v == 0 {
			return false
		}
	}

	return true
}

func canVisitAllRooms1(rooms [][]int) bool {
	visited := make([]int, len(rooms))
	toVisit := make([]int, 1, len(rooms))
	toVisit[0], visited[0] = 0, 0
	start, ending := -1, 0

	for start < ending {
		start += 1
		roomId := toVisit[start]
		for _, nextRoomId := range rooms[roomId] {
			if visited[nextRoomId] == 1 {
				continue
			}
			toVisit = append(toVisit, nextRoomId)
		}
		visited[roomId] = 1
		ending = len(toVisit) - 1
	}

	for _, v := range visited {
		if v == 0 {
			return false
		}
	}

	return true
}
