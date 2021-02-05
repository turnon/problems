/**
 * LC#530 https://leetcode.com/problems/minimum-absolute-difference-in-bst
 * 二叉查找树中任意两结点之差的最小值
 * 思路：中序遍历二叉查找树可得结点从小到大的序列，只要记录序列中相邻结点之差的最小值即可
 */
public class MinimumAbsoluteDifferenceInBST {
    public static void main(String[] args) {
        TreeNode n1 = new TreeNode(
                1,
                null,
                new TreeNode(
                        3,
                        new TreeNode(
                                2,
                                null,
                                null),
                        null)
        );
        TreeNode n2 = new TreeNode(
                1,
                new TreeNode(
                        -1,
                        null,
                        null),
                new TreeNode(
                        30,
                        new TreeNode(
                                20,
                                null,
                                null),
                        null)
        );
        MinimumAbsoluteDifferenceInBST minimumAbsoluteDifferenceInBST = new MinimumAbsoluteDifferenceInBST();
        System.out.println(minimumAbsoluteDifferenceInBST.getMinimumDifference(n1));
        System.out.println(minimumAbsoluteDifferenceInBST.getMinimumDifference(n2));
    }

    public int getMinimumDifference(TreeNode root) {
        Cache cache = new Cache();
        inOrder(root, cache);
        return cache.min;
    }

    private void inOrder(TreeNode node, Cache cache) {
        if (null != node.left) {
            inOrder(node.left, cache);
        }

        if (null != cache.prevInt) {
            int abs = node.val - cache.prevInt;
            if (null == cache.min || abs < cache.min) {
                cache.min = abs;
            }
        }
        cache.prevInt = node.val;

        if (null != node.right) {
            inOrder(node.right, cache);
        }
    }

    class Cache {
        Integer prevInt;
        Integer min;
    }
}
