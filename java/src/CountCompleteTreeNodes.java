/**
 * 222
 * 递归
 */
public class CountCompleteTreeNodes {
    public static void main(String[] args) {
        int c = new CountCompleteTreeNodes().countNodes(new TreeNode(
                1,
                new TreeNode(
                        2,
                        new TreeNode(
                                4,
                                null,
                                null
                        ),
                        new TreeNode(
                                5,
                                null,
                                null
                        )),
                new TreeNode(
                        3,
                        new TreeNode(
                                6,
                                null,
                                null
                        ), null)
        ));

        System.out.println(c);
    }

    public int countNodes(TreeNode root) {
        if (null == root) return 0;
        return 1 + countNodes(root.left) + countNodes(root.right);
    }
}
