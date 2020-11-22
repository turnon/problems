=begin
Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end
@param {TreeNode} p
@param {TreeNode} q
@return {Boolean}

思路：递归比较“左树的左结点和右树的左结点”、“左树的右结点和右树的右结点”

=end
def is_same_tree(p, q)
    return true if p.nil? && q.nil?
    return false if p.nil? || q.nil?

    (p.val == q.val) && is_same_tree(p.left, q.left) && is_same_tree(p.right, q.right)
end
