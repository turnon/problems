import java.util.ArrayList;
import java.util.List;

public class SumRootToLeafNumbers {
    public static void main(String[] args) {
        int sum1 = new SumRootToLeafNumbers().sumNumbers(new TreeNode(
                1,
                new TreeNode(
                        2,
                        null,
                        null),
                new TreeNode(
                        3,
                        null,
                        null)
        ));

        System.out.println(sum1);

        int sum2 = new SumRootToLeafNumbers().sumNumbers(new TreeNode(
                4,
                new TreeNode(
                        9,
                        new TreeNode(
                                5,
                                null,
                                null
                        ),
                        new TreeNode(
                                1,
                                null,
                                null
                        )),
                new TreeNode(
                        0,
                        null,
                        null)
        ));
        System.out.println(sum2);
    }

    public int sumNumbers(TreeNode root) {
        if (null == root) {
            return 0;
        }
        Result result = new Result();
        sumNumbers(root, new ArrayList<Integer>(), result);
        return result.result;
    }

    public void sumNumbers(TreeNode root, List<Integer> path, Result result) {
        path.add(root.val);
        if (null == root.left && null == root.right) {
            result.result = result.result + sumPath(path);
        }
        if (null != root.left) {
            sumNumbers(root.left, path, result);
        }
        if (null != root.right) {
            sumNumbers(root.right, path, result);
        }
        path.remove(path.size() - 1);
    }

    private int sumPath(List<Integer> path) {
        int size = path.size();
        Integer sum = 0;
        for (int i = 0; i < size; i++) {
            Integer val = path.get(size - i - 1);
            int times = (int) Math.pow(10, i);
            sum = sum + (val * times);
        }
        return sum;
    }

    class Result {
        Integer result = 0;
    }

//    public int sumNumbers(TreeNode root) {
//        return sum(root, 0);
//    }
//
//    public int sum(TreeNode n, int s) {
//        if (n == null) return 0;
//        int sumUntilThisNode = s * 10 + n.val;
//        if (n.right == null && n.left == null) return sumUntilThisNode;
//        return sum(n.left, sumUntilThisNode) + sum(n.right, sumUntilThisNode);
//    }
}
