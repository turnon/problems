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
@param {TreeNode} root
@return {Integer[][]}

思路：层序遍历树，并使用enumerator返回每一层，对每层顺序或反序遍历

=end
def zigzag_level_order(root)
    return [] unless root

    current_level = [root]
    node_levels = Enumerator.new do |yielder|
        until current_level.size == 0
            yielder << current_level
            current_level = current_level.each_with_object([]) do |n, arr|
                arr << n.left if n.left
                arr << n.right if n.right
            end
        end
    end

    reverse = false
    value_levels = []
    node_levels.each do |nodes|
        values = []
        nodes.send(reverse ? :reverse_each : :each) do |n|
            values << n.val
        end
        reverse = !reverse
        value_levels << values if values.count > 0
    end
    value_levels
end
