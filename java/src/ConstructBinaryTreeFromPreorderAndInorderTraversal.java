//序列的第一个元素为根结点，用其分拆中序列可得左右子树的中序列。先序列尾部的前段为左子树的先序列，后段为右子树的先序列，可用刚才分出的两个中序列的长度来拆出这两个先序列。递归可得各层子树的根节点。使用类Slice抽象分段，辅助思考。


public class ConstructBinaryTreeFromPreorderAndInorderTraversal {
    public static void main(String[] args) {
        int[] preorder = new int[]{3, 9, 20, 15, 7};
        int[] inorder = new int[]{9, 3, 15, 20, 7};
        TreeNode node = new ConstructBinaryTreeFromPreorderAndInorderTraversal().buildTree(preorder, inorder);
        System.out.println(node.preorder());
        System.out.println(node.inorder());
    }

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return buildTree(
                new Slice(preorder, 0, preorder.length - 1),
                new Slice(inorder, 0, inorder.length - 1)
        );
    }

    public TreeNode buildTree(Slice preorder, Slice inorder) {
        if (preorder.size() == 0) return null;

        TreeNode root = preorder.first();
        Slice[] left_right_inorder = inorder.split(root);

        Slice left_inorder = left_right_inorder[0];
        if (left_inorder.size() > 0) {
            root.left = buildTree(
                    preorder.cut(1, 1 + left_inorder.size()),
                    left_inorder
            );
        }

        Slice right_inorder = left_right_inorder[1];
        if (right_inorder.size() > 0) {
            root.right = buildTree(
                    preorder.cut(left_inorder.size() + 1, left_inorder.size() + 1 + right_inorder.size()),
                    right_inorder
            );
        }

        return root;
    }

    class Slice {
        int[] array;
        int start;
        int end;

        Slice(int[] array, int start, int end) {
            this.array = array;
            this.start = start;
            this.end = end;
        }

        TreeNode first() {
            return start > end ? null : new TreeNode(array[start]);
        }

        Slice[] split(TreeNode node) {
            int mid = start;
            while (mid < end) {
                if (array[mid] == node.val) {
                    break;
                }
                mid++;
            }
            Slice left = new Slice(array, start, mid - 1);
            Slice right = new Slice(array, mid + 1, end);
            return new Slice[]{left, right};
        }

        Slice cut(int from, int len) {
            return new Slice(array, start + from, start + from + len);
        }

        int size() {
            return end - start + 1;
        }
    }
}
