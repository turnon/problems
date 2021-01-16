import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

/**
 * 257 https://leetcode.com/problems/binary-tree-paths/
 * 所有根到叶子的路径，用字符串表示
 * 思路：递归，左右叶子结点返回仅含自己的路径数组，本节点再把自己追加上去，顶层将所有数组反转并拼成字符串
 */
public class BinaryTreePaths {
    public static void main(String[] args) {
        List<String> paths = new BinaryTreePaths().binaryTreePaths(new TreeNode(
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
                        null)
        ));

        System.out.println(paths);
    }

    public List<String> binaryTreePaths(TreeNode root) {
        return null == root ? List.of() : binaryTreePathBuilders(root).stream().map(path -> {
            Collections.reverse(path);
            List<String> pathStr = path.stream().map(i -> i.toString()).collect(Collectors.toList());
            return String.join("->", pathStr);
        }).collect(Collectors.toList());
    }

    private List<List<Integer>> binaryTreePathBuilders(TreeNode root) {
        if (null == root.left && null == root.right) {
            List<List<Integer>> paths = new ArrayList<>();
            List<Integer> path = new ArrayList<>();
            path.add(root.val);
            paths.add(path);
            return paths;
        }
        List<List<Integer>> paths = new ArrayList<>();
        if (null != root.left) {
            paths.addAll(binaryTreePathBuilders(root.left));
        }
        if (null != root.right) {
            paths.addAll(binaryTreePathBuilders(root.right));
        }
        for (List<Integer> path : paths) {
            path.add(root.val);
        }
        return paths;
    }
}
