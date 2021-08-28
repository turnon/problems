package main

import "testing"

func TestNetworkDelayTime1(t *testing.T) {
	res := networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2)
	testNetworkDelayTime(res, 2, t)
}

func TestNetworkDelayTime2(t *testing.T) {
	res := networkDelayTime([][]int{{1, 2, 1}}, 2, 1)
	testNetworkDelayTime(res, 1, t)
}

func TestNetworkDelayTime3(t *testing.T) {
	res := networkDelayTime([][]int{{1, 2, 1}}, 2, 2)
	testNetworkDelayTime(res, -1, t)
}

func testNetworkDelayTime(act, exp int, t *testing.T) {
	if act != exp {
		t.Error(act, "should be", exp)
	}
}
