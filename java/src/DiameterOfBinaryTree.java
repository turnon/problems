/**
 * LC#543 https://leetcode.com/problems/diameter-of-binary-tree/
 * 二叉树中最远的两点距离
 * 思路：左右子树的深度相加，但要提防左右不太对称的情况，所以还要记住已计算过的最大距离
 */
public class DiameterOfBinaryTree {
    public static void main(String[] args) {
        TreeNode n1 = new TreeNode(
                1,
                new TreeNode(
                        2,
                        new TreeNode(
                                4,
                                null,
                                null),
                        new TreeNode(
                                5,
                                null,
                                null)),
                new TreeNode(3, null, null)

        );

        DiameterOfBinaryTree diameterOfBinaryTree = new DiameterOfBinaryTree();
        int d = diameterOfBinaryTree.diameterOfBinaryTree(n1);
        System.out.println(d);

    }

    public int diameterOfBinaryTree(TreeNode root) {
        MaxDiameter maxDiameter = new MaxDiameter();
        maxDiameter.setDiameter(null == root ? 0 : depth(root.left, maxDiameter) + depth(root.right, maxDiameter));
        return maxDiameter.value;
    }

    private int depth(TreeNode root, MaxDiameter maxDiameter) {
        if (null == root) {
            return 0;
        }
        int leftDepth = depth(root.left, maxDiameter);
        int rightDepth = depth(root.right, maxDiameter);
        maxDiameter.setDiameter(leftDepth + rightDepth);
        return 1 + Math.max(leftDepth, rightDepth);
    }

    class MaxDiameter {
        Integer value;

        void setDiameter(Integer diameter) {
            if (null == value || value < diameter) {
                value = diameter;
            }
        }
    }
}
