package main

/*
LC#654 Maximum Binary Tree
思路：取数组最大值索引分割出左右数组作为左右子树，然后递归处理左右数组
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	var rootIdx int
	var rootVal int
	for idx, val := range nums {
		if val > rootVal {
			rootIdx = idx
			rootVal = val
		}
	}

	left := nums[0:rootIdx]
	var right []int
	if rootIdx == length-1 {
		right = []int{}
	} else {
		right = nums[rootIdx+1:]
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  constructMaximumBinaryTree(left),
		Right: constructMaximumBinaryTree(right),
	}
}
