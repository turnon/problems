/*
 * 一个将树转化为链表并返回尾结点的递归函数，先处理左树再处理右树，并用左树的尾结点指向右树
 */
public class FlattenBinaryTreeToLinkedList {
    public static void main(String[] args) {
        TreeNode n = new TreeNode(
                1,
                new TreeNode(
                        2,
                        new TreeNode(
                                3,
                                null,
                                null
                        ),
                        new TreeNode(
                                4,
                                null,
                                null
                        )),
                new TreeNode(
                        5,
                        null,
                        new TreeNode(
                                6,
                                null,
                                null
                        ))
        );
        new FlattenBinaryTreeToLinkedList().flatten(n);

        System.out.println(n.inorder());
    }

    public void flatten(TreeNode root) {
        if (null != root) {
            toLinkedList(root);
        }
    }

    private TreeNode toLinkedList(TreeNode root) {
        TreeNode left = root.left;
        TreeNode leftTail = null;
        if (null == left) {
            leftTail = root;
        } else {
            leftTail = toLinkedList(left);
            root.left = null;
        }

        TreeNode right = root.right;
        root.right = left;
        if (null == right) {
            return leftTail;
        } else {
            leftTail.right = right;
            return toLinkedList(right);
        }
    }

//    /*
//     * 论坛解法
//     */
//    public void flatten(TreeNode root) {
//        flatten(root, null);
//    }
//
//    private TreeNode flatten(TreeNode root, TreeNode pre) {
//        if (root == null) return pre;
//        pre = flatten(root.right, pre);
//        pre = flatten(root.left, pre);
//        root.right = pre;
//        root.left = null;
//        pre = root;
//        return pre;
//    }
}
