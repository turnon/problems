/**
 * LC#563 https://leetcode.com/problems/binary-tree-tilt/
 * 各结点倾斜度之和
 * 思路：递归后序遍历，并传入一个辅助对象累加每个结点的倾斜度
 */
public class BinaryTreeTilt {
    public static void main(String[] args) {
        TreeNode n1 = new TreeNode(
                1,
                new TreeNode(2, null, null),
                new TreeNode(3, null, null)
        );

        TreeNode n2 = new TreeNode(
                4,
                new TreeNode(2,
                        new TreeNode(3, null, null),
                        new TreeNode(5, null, null)),
                new TreeNode(9,
                        null,
                        new TreeNode(7, null, null))
        );

        BinaryTreeTilt binaryTreeTilt = new BinaryTreeTilt();
        System.out.println(binaryTreeTilt.findTilt(n1));
        System.out.println(binaryTreeTilt.findTilt(n2));
    }

    public int findTilt(TreeNode root) {
        if (null == root) {
            return 0;
        }
        Tilt tilt = new Tilt();
        sum(root, tilt);
        return tilt.val;
    }

    private int sum(TreeNode root, Tilt tilt) {
        int leftSum = (null == root.left ? 0 : sum(root.left, tilt));
        int rightSum = (null == root.right ? 0 : sum(root.right, tilt));
        tilt.val += Math.abs(leftSum - rightSum);
        return root.val + leftSum + rightSum;
    }

    class Tilt {
        int val;
    }
}
