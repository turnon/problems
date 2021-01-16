/**
 * 404 https://leetcode.com/problems/sum-of-left-leaves/
 * 所有左叶子结点的和
 * 思路：递归，并告诉下层他是不是左结点，他若是叶子结点并被告知是左结点，则返回自身的值，否则返回零
 */
public class SumOfLeftLeaves {
    public static void main(String[] args) {
        int sum = new SumOfLeftLeaves().sumOfLeftLeaves(new TreeNode(
                3,
                new TreeNode(
                        9,
                        null,
                        null),
                new TreeNode(
                        20,
                        new TreeNode(
                                15,
                                null,
                                null),
                        new TreeNode(
                                7,
                                null,
                                null))
        ));

        System.out.println(sum);
    }

    public int sumOfLeftLeaves(TreeNode root) {
        return null == root ? 0 : sumOfLeftLeaves(root, false);
    }

    private int sumOfLeftLeaves(TreeNode root, boolean isLeft) {
        if (null == root.left && null == root.right) {
            return isLeft ? root.val : 0;
        }
        return (null == root.left ? 0 : sumOfLeftLeaves(root.left, true)) +
                (null == root.right ? 0 : sumOfLeftLeaves(root.right, false));
    }
}
