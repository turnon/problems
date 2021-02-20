/**
 * LC#572 https://leetcode.com/problems/subtree-of-another-tree/
 * 判断是否子树
 * 思路：递归检查是否当前结点相等且两树的左右结点也分别相等，需要增加一个标识变量指示 t 结点是否根结点。
 * 论坛答案则是以树为单位作对比，不需要额外的变量，貌似因此而更高效
 */
public class SubtreeOfAnotherTree {
    public static void main(String[] args) {
        TreeNode n0 = new TreeNode(
                4,
                new TreeNode(1, null, null),
                new TreeNode(2, null, null)
        );

        TreeNode n1 = new TreeNode(
                3,
                new TreeNode(4,
                        new TreeNode(1, null, null),
                        new TreeNode(2, null, null)),
                new TreeNode(5,
                        null,
                        null)
        );

        TreeNode n2 = new TreeNode(
                3,
                new TreeNode(4,
                        new TreeNode(1, null, null),
                        new TreeNode(2,
                                new TreeNode(0, null, null)
                                , null)),
                new TreeNode(5,
                        null,
                        null)
        );

        TreeNode m0 = new TreeNode(
                3,
                new TreeNode(1,
                        null,
                        null
                ),
                new TreeNode(2,
                        null,
                        null)
        );

        TreeNode m1 = new TreeNode(
                3,
                new TreeNode(4,
                        new TreeNode(1, null, null),
                        null
                ),
                new TreeNode(5,
                        new TreeNode(2, null, null),
                        null)
        );

        SubtreeOfAnotherTree subtreeOfAnotherTree = new SubtreeOfAnotherTree();
        System.out.println(subtreeOfAnotherTree.isSubtree(n1, n0));
        System.out.println(subtreeOfAnotherTree.isSubtree(n2, n0));
        System.out.println(subtreeOfAnotherTree.isSubtree(m1, m0));
    }

    public boolean isSubtree(TreeNode s, TreeNode t) {
        return isSubtree(s, t, true);
    }

    private boolean isSubtree(TreeNode s, TreeNode t, boolean tIsRoot) {
        if (null == s && null == t) {
            return true;
        }
        if (null == s || null == t) {
            return false;
        }
        if (
                (tIsRoot && isSubtree(s.left, t, true)) ||
                        (tIsRoot && isSubtree(s.right, t, true)) ||
                        (s.val == t.val && isSubtree(s.left, t.left, false) && isSubtree(s.right, t.right, false))
        ) {
            return true;
        }
        return false;
    }
}
