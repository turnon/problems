/**
 * 230 二叉树中第K小的结点
 * 中序遍历，并传入一个辅助对象记录当前结点的排位，如果排位等于K，则记下当前结点，停止遍历。
 */
public class KthSmallestElementInABST {
    public static void main(String[] args) {
        int value1 = new KthSmallestElementInABST().kthSmallest(new TreeNode(
                3,
                new TreeNode(
                        1,
                        null,
                        new TreeNode(
                                2,
                                null,
                                null
                        )),
                new TreeNode(
                        4,
                        null,
                        null)
        ), 1);

        int value2 = new KthSmallestElementInABST().kthSmallest(new TreeNode(
                5,
                new TreeNode(
                        3,
                        new TreeNode(
                                2,
                                new TreeNode(
                                        1,
                                        null,
                                        null
                                ),
                                null
                        ),
                        new TreeNode(
                                4,
                                null,
                                null
                        )),
                new TreeNode(
                        6,
                        null,
                        null)
        ), 3);

        System.out.println(value1);
        System.out.println(value2);
    }

    public int kthSmallest(TreeNode root, int k) {
        Context ctx = new Context(k);
        inOrder(root, ctx);
        return ctx.node.val;
    }

    private void inOrder(TreeNode root, Context ctx) {
        if (null == root || null != ctx.node) {
            return;
        }

        inOrder(root.left, ctx);

        if (null != ctx.node) {
            return;
        }

        if (ctx.wantedK == ++ctx.currentK) {
            ctx.node = root;
            return;
        }

        inOrder(root.right, ctx);
    }

    class Context {
        private int wantedK;
        private int currentK;
        private TreeNode node;

        public Context(int wantedK) {
            this.wantedK = wantedK;
            this.currentK = 0;
        }
    }
}
