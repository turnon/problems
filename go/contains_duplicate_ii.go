package main

/*
LC#219 在K步内是否有重复数字
https://leetcode.com/problems/contains-duplicate-ii
思路：滑窗。（哈希表直接记录每个数的的索引也可以，但若数组很大且不存在重复，则哈希表也会变得很大）
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	if k < 1 || len(nums) <= 1 {
		return false
	}

	hash := make(map[int]bool)
	add := func(num int) (dup bool) {
		if hash[num] {
			return true
		}
		hash[num] = true
		return false
	}
	rm := func(num int) {
		hash[num] = false
	}

	for i := 0; i <= k && i < length; i++ {
		if add(nums[i]) {
			return true
		}
	}

	windowStart := 0
	for windowEnd := k + 1; windowEnd < length; windowEnd++ {
		rm(nums[windowStart])
		if add(nums[windowEnd]) {
			return true
		}
		windowStart++
	}
	return false
}
