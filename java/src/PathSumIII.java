import java.util.ArrayList;
import java.util.List;

/*
 * 思路：
 * 递归深度遍历，并用一个栈记录所经过的结点，
 * 进入新结点时把新结点的值加到栈上的每一帧，可得每段与本结点相连的路径的和，
 * 如果和等于目标值，则计数加一。
 * 退出结点时需恢复栈上的和。
 */
public class PathSumIII {
    public static void main(String[] args) {
        int count = new PathSumIII().pathSum(new TreeNode(
                10,
                new TreeNode(
                        5,
                        new TreeNode(
                                3,
                                new TreeNode(
                                        3,
                                        null,
                                        null
                                ),
                                new TreeNode(
                                        -2,
                                        null,
                                        null
                                )
                        ),
                        new TreeNode(
                                2,
                                null,
                                new TreeNode(
                                        1,
                                        null,
                                        null
                                )
                        )),
                new TreeNode(
                        -3,
                        null,
                        new TreeNode(
                                11,
                                null,
                                null
                        ))
        ), 8);

        System.out.println(count);
    }

    public int pathSum(TreeNode root, int sum) {
        if (null == root) {
            return 0;
        }
        return pathSum(root, sum, new ArrayList<Integer>());
    }

    private int pathSum(TreeNode root, Integer target, List<Integer> stack) {
        int count = 0;
        int val = root.val;

        for (int i = 0; i < stack.size(); i++) {
            Integer sum = stack.get(i) + val;
            if (sum.equals(target)) {
                System.out.println(stack);
                System.out.println(stack.get(i));
                System.out.println(val);
                System.out.println("------");
                count++;
            }
            stack.set(i, sum);
        }

        if (target.equals(val)) {
            count++;
        }

        stack.add(val);
        if (null != root.left) {
            count += pathSum(root.left, target, stack);
        }
        if (null != root.right) {
            count += pathSum(root.right, target, stack);
        }
        stack.remove(stack.size() - 1);

        for (int i = 0; i < stack.size(); i++) {
            Integer recover = stack.get(i) - val;
            stack.set(i, recover);
        }

        return count;
    }

//    /*
//     * 论坛上十分简洁的解法……
//     */
//    public int pathSum(TreeNode root, int sum) {
//        if (root == null) return 0;
//        return pathSumFrom(root, sum) + pathSum(root.left, sum) + pathSum(root.right, sum);
//    }
//
//    private int pathSumFrom(TreeNode node, int sum) {
//        if (node == null) return 0;
//        return (node.val == sum ? 1 : 0)
//                + pathSumFrom(node.left, sum - node.val) + pathSumFrom(node.right, sum - node.val);
//    }

}
