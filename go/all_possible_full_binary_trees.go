package main

/*
LC#894 所有满二叉树组合
https://leetcode.com/problems/all-possible-full-binary-trees
思路：递归生成所有子树的满二叉树组合；右子树可利用左子树的缓存。
*/
func allPossibleFBT(n int) []*TreeNode {
	cache := make(map[int][]*TreeNode)

	var _allPossibleFBT func(n int) []*TreeNode
	_allPossibleFBT = func(n int) []*TreeNode {
		if possibles, ok := cache[n]; ok {
			return possibles
		}

		childrenCount := n - 1
		if childrenCount == 0 {
			return []*TreeNode{&TreeNode{0, nil, nil}}
		}

		result := []*TreeNode{}
		for leftChildrenCount := 0; leftChildrenCount <= childrenCount; leftChildrenCount++ {
			rightChildrenCount := childrenCount - leftChildrenCount
			if leftChildrenCount%2 == 0 || rightChildrenCount%2 == 0 {
				continue
			}
			leftPossibles := _allPossibleFBT(leftChildrenCount)
			righttPossibles := _allPossibleFBT(rightChildrenCount)
			for _, lp := range leftPossibles {
				for _, rp := range righttPossibles {
					result = append(result, &TreeNode{0, lp, rp})
				}
			}
		}

		cache[n] = result
		return result
	}

	return _allPossibleFBT(n)
}
