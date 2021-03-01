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
@return {Boolean}

思路：层序遍历，使用enumerator返回每层，再判断每层是否回文。
（也可使用链表作队列层序遍历，每次出队两个结点来比较，
再按顺序入队该两结点n1左结点、n2左结点、n1右结点、n2右结点，
奈何ruby没有自带链表）

=end
def is_symmetric(root)
    level = [root]

    levels = Enumerator.new do |yielder|
        until level.all?(&:nil?)
            ret = level
            level = level.each_with_object([]) do |node, arr|
                if node.nil?
                    arr << nil << nil
                else
                    arr << node.left << node.right
                end
            end
            yielder << ret
        end
    end

    levels.each do |level|
        mid = level.count / 2
        level.each_with_index do |node, i|
            break if i >= mid
            l = level[i]
            r = level[-i - 1]
            next if (l.nil? && r.nil?) || (!l.nil? && !r.nil? && l.val == r.val)
            return false
        end
    end

    true
end
