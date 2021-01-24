/**
 * 450 https://leetcode.com/problems/delete-node-in-a-bst/
 * 删除二叉查找树中的某个结点
 * 思路：二分查找出结点所在位置，并记录前一结点，
 * 让前一结点的子树指向被删节点的左树（或右树），
 * 被删结点的右树（或左树）接到左树的最大结点下（或右树的最小结点下）。
 * 需留意被删的是否就是根节点。
 */
public class DeleteNodeInBST {
    public static void main(String[] args) {
        DeleteNodeInBST d = new DeleteNodeInBST();

        TreeNode n1 = d.deleteNode(new TreeNode(
                5,
                new TreeNode(
                        3,
                        new TreeNode(2, null, null),
                        new TreeNode(4, null, null)),
                new TreeNode(
                        6,
                        null,
                        new TreeNode(7, null, null))
        ), 3);

        System.out.println(n1.inorder());

        TreeNode n2 = d.deleteNode(new TreeNode(
                5,
                new TreeNode(
                        3,
                        new TreeNode(2, null, null),
                        new TreeNode(4, null, null)),
                new TreeNode(
                        6,
                        null,
                        new TreeNode(7, null, null))
        ), 0);

        System.out.println(n2.inorder());

        TreeNode n3 = d.deleteNode(new TreeNode(
                5,
                new TreeNode(
                        3,
                        new TreeNode(2, null, null),
                        new TreeNode(4, null, null)),
                new TreeNode(
                        6,
                        null,
                        new TreeNode(7, null, null))
        ), 5);

        System.out.println(n3.inorder());

        TreeNode n4 = d.deleteNode(new TreeNode(
                3,
                new TreeNode(
                        1,
                        null,
                        new TreeNode(2, null, null)),
                new TreeNode(
                        4,
                        null,
                        null)
        ), 2);

        System.out.println(n4.inorder());
    }

    public TreeNode deleteNode(TreeNode root, int key) {
        TreeNode curr = root;
        TreeNode prev = null;
        while (null != curr && key != curr.val) {
            prev = curr;
            curr = key < curr.val ? curr.left : curr.right;
        }
        if (null != curr) {
            // 要刪的是root
            if (null == prev) {
                if (null != root.left && null != root.right) {
                    max(root.left).right = root.right;
                    root = root.left;
                } else {
                    root = null == curr.left ? curr.right : curr.left;
                }
            } else {
                if (null != curr.left && null != curr.right) {
                    max(curr.left).right = curr.right;
                    if (prev.val > curr.val) {
                        prev.left = curr.left;
                    } else {
                        prev.right = curr.left;
                    }
                } else {
                    TreeNode next = null == curr.left ? curr.right : curr.left;
                    if (prev.val > curr.val) {
                        prev.left = next;
                    } else {
                        prev.right = next;
                    }
                }
            }
        }
        return root;
    }

    private TreeNode max(TreeNode node) {
        while (null != node.right) {
            node = node.right;
        }
        return node;
    }
}
