public class PathSum {
    public static void main(String[] args) {
        boolean bool = new PathSum().hasPathSum(new TreeNode(
                5,
                new TreeNode(
                        4,
                        new TreeNode(
                                11,
                                new TreeNode(
                                        7,
                                        null,
                                        null
                                ),
                                new TreeNode(
                                        2,
                                        null,
                                        null
                                )
                        ),
                        null),
                new TreeNode(
                        8,
                        new TreeNode(
                                13,
                                null,
                                null
                        ),
                        new TreeNode(
                                4,
                                null,
                                new TreeNode(
                                        1,
                                        null,
                                        null
                                )
                        ))
        ), 22);

        System.out.println(bool);
    }

    public boolean hasPathSum(TreeNode root, int sum) {
        if (null == root) {
            return false;
        }

        if (root.val == sum && root.left == null && root.right == null) {
            return true;
        }

        int rest = sum - root.val;
        return hasPathSum(root.left, rest) || hasPathSum(root.right, rest);
    }
}
