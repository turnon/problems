package main

/*
LC#341 嵌套列表迭代器
https://leetcode.com/problems/flatten-nested-list-iterator
思路：创建队列保存嵌套列表。
每次Next都要检查队头是否嵌套列表，若是则提取与后尾拼接为新队列，直至队头不为嵌套列表，再出队返回。
而HasNext就是递归检查队列中是否存在不为嵌套列表的元素
*/
type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
	q []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{q: nestedList}
}

func (this *NestedIterator) Next() int {
	for {
		head, tail := this.q[0], this.q[1:]
		if head.IsInteger() {
			this.q = tail
			return head.GetInteger()
		}

		this.q = append(head.GetList(), tail...)
	}
}

func (this *NestedIterator) HasNext() bool {
	if len(this.q) == 0 {
		return false
	}
	for _, nestedInt := range this.q {
		if nestedInt.IsInteger() || Constructor(nestedInt.GetList()).HasNext() {
			return true
		}
	}
	return false
}
