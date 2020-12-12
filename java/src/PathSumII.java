import java.util.ArrayList;
import java.util.List;

public class PathSumII {
    public static void main(String[] args) {
        List<List<Integer>> arr = new PathSumII().pathSum(new TreeNode(
                5,
                new TreeNode(
                        4,
                        new TreeNode(
                                11,
                                new TreeNode(
                                        7,
                                        null,
                                        null
                                ),
                                new TreeNode(
                                        2,
                                        null,
                                        null
                                )
                        ),
                        null),
                new TreeNode(
                        8,
                        new TreeNode(
                                13,
                                null,
                                null
                        ),
                        new TreeNode(
                                4,
                                new TreeNode(
                                        5,
                                        null,
                                        null
                                ),
                                new TreeNode(
                                        1,
                                        null,
                                        null
                                )
                        ))
        ), 22);

        System.out.println(arr);
    }

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        if (null == root) {
            return result;
        }
        List<Integer> stack = new ArrayList<Integer>();

        collectPathSum(root, sum, stack, result);
        return result;
    }

    private void collectPathSum(TreeNode root, int sum, List<Integer> stack, List<List<Integer>> result) {
        stack.add(root.val);

        if (root.left == null && root.right == null) {
            if (root.val == sum) {
                List<Integer> subResult = new ArrayList<Integer>(stack.size());
                subResult.addAll(stack);
                result.add(subResult);
            }
            pop(stack);
            return;
        }

        int remain = sum - root.val;
        if (root.left != null) {
            collectPathSum(root.left, remain, stack, result);
        }
        if (root.right != null) {
            collectPathSum(root.right, remain, stack, result);
        }
        pop(stack);
    }

    private void pop(List<Integer> stack) {
        stack.remove(stack.size() - 1);
    }
}
