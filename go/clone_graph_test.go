package main

import "testing"

func TestCloneGraph1(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	m1 := cloneGraph(n1)
	_testCloneGraph(n1, t)
	t.Log("------------")
	_testCloneGraph(m1, t)
}

func _testCloneGraph(n *Node, t *testing.T) {
	printed := make(map[*Node]bool)
	var _printG func(n *Node)
	_printG = func(n *Node) {
		if _, ok := printed[n]; ok {
			return
		}
		printed[n] = true
		t.Log(n.Val)
		for _, nei := range n.Neighbors {
			_printG(nei)
		}
	}
	_printG(n)
}
