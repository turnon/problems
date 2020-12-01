import java.util.ArrayList;

public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }

    ArrayList preorder() {
        return preorder(new ArrayList());
    }

    ArrayList preorder(ArrayList order) {
        order.add(val);
        if (null != left) order.addAll(left.preorder());
        if (null != right) order.addAll(right.preorder());
        return order;
    }

    ArrayList inorder() {
        return inorder(new ArrayList());
    }

    ArrayList inorder(ArrayList order) {
        if (null != left) order.addAll(left.inorder());
        order.add(val);
        if (null != right) order.addAll(right.inorder());
        return order;
    }
}
