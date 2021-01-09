import java.util.ArrayList;
import java.util.List;

/**
 * 226 反转二叉树
 * 递归或栈
 */
public class InvertBinaryTree {
    public static void main(String[] args) {
        TreeNode node = new InvertBinaryTree().invertTree(new TreeNode(
                4,
                new TreeNode(
                        2,
                        new TreeNode(
                                1,
                                null,
                                null
                        ),
                        new TreeNode(
                                3,
                                null,
                                null
                        )),
                new TreeNode(
                        7,
                        new TreeNode(
                                6,
                                null,
                                null
                        ),
                        new TreeNode(
                                9,
                                null,
                                null
                        ))
        ));

        System.out.println(node.preorder());
    }

    public TreeNode invertTree(TreeNode root) {
        if (null == root) return root;
        TreeNode tmp = invertTree(root.left);
        root.left = invertTree(root.right);
        root.right = tmp;
        return root;
    }

//    public TreeNode invertTree(TreeNode root) {
//        if (null == root) return root;
//
//        List<TreeNode> stack = new ArrayList<TreeNode>();
//        stack.add(root);
//        int size = stack.size();
//
//        while (size > 0) {
//            TreeNode node = stack.remove(size - 1);
//            TreeNode tmp = node.left;
//            node.left = node.right;
//            node.right = tmp;
//            if (null != node.left) stack.add(node.left);
//            if (null != node.right) stack.add(node.right);
//            size = stack.size();
//        }
//        return root;
//    }
}
