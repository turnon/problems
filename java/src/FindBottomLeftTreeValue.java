import java.util.LinkedList;

/**
 * LC#513 https://leetcode.com/problems/find-bottom-left-tree-value/
 * 找出底层最左的结点
 * 思路：用队列层序遍历，记录前一出队结点，并在每层结束时插入一个标记结点，
 * 如果标记结点出队后队列全空，则前一出队结点为底层最左的结点
 */
public class FindBottomLeftTreeValue {
    public static void main(String[] args) {
        TreeNode n1 = new TreeNode(
                2,
                new TreeNode(
                        1,
                        null,
                        null),
                new TreeNode(
                        3,
                        null,
                        null)
        );

        TreeNode n2 = new TreeNode(
                1,
                new TreeNode(
                        2,
                        new TreeNode(
                                4,
                                null,
                                null
                        ),
                        null),
                new TreeNode(
                        3,
                        new TreeNode(
                                5,
                                new TreeNode(
                                        7,
                                        null,
                                        null
                                ),
                                null
                        ),
                        new TreeNode(
                                6,
                                null,
                                null
                        ))
        );
        FindBottomLeftTreeValue findBottomLeftTreeValue = new FindBottomLeftTreeValue();
        System.out.println(findBottomLeftTreeValue.findBottomLeftValue(n1));
        System.out.println(findBottomLeftTreeValue.findBottomLeftValue(n2));
    }

    public int findBottomLeftValue(TreeNode root) {
        LinkedList<TreeNode> q = new LinkedList<>();
        TreeNode levelEnd = new TreeNode(0);
        TreeNode prev = root;
        TreeNode leftmost = root;
        q.add(root);
        q.add(levelEnd);

        while (!q.isEmpty()) {
            TreeNode curr = q.removeFirst();
            if (levelEnd == curr) {
                if (q.isEmpty()) {
                    break;
                } else {
                    q.add(curr);
                }
            } else {
                if (null != curr.left) {
                    q.add(curr.left);
                }
                if (null != curr.right) {
                    q.add(curr.right);
                }
                if (levelEnd == prev) {
                    leftmost = curr;
                }
            }
            prev = curr;
        }

        return leftmost.val;
    }
}
