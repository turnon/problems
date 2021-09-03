package main

/*
LC#997 寻找判官
https://leetcode.com/problems/find-the-town-judge
思路：记录每个人被信任的票数，但如果他有信任别人，就将其被信任票数扣成负数。最后检查哪个人票数等于其他人的总数。
*/
func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return 1
	}

	trustedCount := make([]int, n+1)

	for _, trustPair := range trust {
		trusting, trusted := trustPair[0], trustPair[1]
		trustedCount[trusted] = trustedCount[trusted] + 1
		trustedCount[trusting] = trustedCount[trusting] - n
	}

	othersCount := n - 1
	for trusted, count := range trustedCount {
		if count == othersCount {
			return trusted
		}
	}
	return -1
}
