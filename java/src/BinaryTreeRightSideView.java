import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * LC#199
 * 使用链表作队列来层序遍历，并在每层的结尾插入一个假结点（可重复使用），
 * 假结点出队时，前一结点就是该层右视图所看到的结点
 */
public class BinaryTreeRightSideView {
    public static void main(String[] args) {
        List<Integer> list = new BinaryTreeRightSideView().rightSideView(new TreeNode(
                1,
                new TreeNode(
                        2,
                        null,
                        new TreeNode(
                                5,
                                null,
                                null
                        )),
                new TreeNode(
                        3,
                        null,
                        new TreeNode(
                                4,
                                null,
                                null
                        ))
        ));

        System.out.println(list);
    }

    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> list = new ArrayList<>();
        if (null == root) {
            return list;
        }

        TreeNode prev = null;
        TreeNode curr = null;
        TreeNode levelEnd = new TreeNode();
        LinkedList<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        queue.add(levelEnd);

        while (!queue.isEmpty()) {
            curr = queue.removeFirst();
            if (curr == levelEnd) {
                list.add(prev.val);
                if (!queue.isEmpty()) {
                    queue.add(curr);
                }
            } else {
                if (null != curr.left) {
                    queue.add(curr.left);
                }
                if (null != curr.right) {
                    queue.add(curr.right);
                }
                prev = curr;
            }
        }

        return list;
    }
}
