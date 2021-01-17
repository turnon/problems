/**
 * 297 https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
 * 自定义树的序列化和反序列化
 * 思路：自定序列化格式为“结点值+逗号+左子树长度+（逗号+序列化的左子树）+（逗号+序列化的右子树）”，如果子树为空，则括号内容为空字符串。
 * 反序列化时根据逗号分隔读取结点值和左子树长度，可计算出左右子树偏移量，然后递归处理左右子树
 */
public class SerializeAndDeserializeBinaryTree {
    public static void main(String[] args) {
        SerializeAndDeserializeBinaryTree sd = new SerializeAndDeserializeBinaryTree();
        String str1 = sd.serialize(new TreeNode(
                1,
                new TreeNode(
                        2,
                        null,
                        null),
                new TreeNode(
                        3,
                        new TreeNode(
                                4,
                                null,
                                null),
                        new TreeNode(
                                5,
                                null,
                                null))
        ));

        System.out.println(str1);
        System.out.println(sd.deserialize(str1).preorder());

        String str2 = sd.serialize(null);
        System.out.println(str2);
        TreeNode node2 = sd.deserialize(str2);
        System.out.println(null == node2 ? "null" : node2.preorder());

        String str3 = sd.serialize(new TreeNode(1, null, null));
        System.out.println(str3);
        System.out.println(sd.deserialize(str3).preorder());

        String str4 = sd.serialize(new TreeNode(1, new TreeNode(2, null, null), null));
        System.out.println(str4);
        System.out.println(sd.deserialize(str4).preorder());

        String str5 = sd.serialize(new TreeNode(1, null, new TreeNode(2, null, null)));
        System.out.println(str5);
        System.out.println(sd.deserialize(str5).preorder());
    }

    public String serialize(TreeNode root) {
        if (null == root) return "";
        String left = serialize(root.left);
        String right = serialize(root.right);
        StringBuilder sb = new StringBuilder();
        sb.append(root.val).append(",").append(left.length());
        if (left.length() != 0) sb.append(",").append(left);
        if (right.length() != 0) sb.append(",").append(right);
        return sb.toString();
    }

    public TreeNode deserialize(String data) {
        return "".equals(data) ? null : deserialize(data, 0, data.length()-1);
    }

    private TreeNode deserialize(String data, int start, int end) {
        int valEnd = data.indexOf(",", start);
        int val = Integer.parseInt(data.substring(start, valEnd));
        TreeNode node = new TreeNode(val);

        valEnd = valEnd + 1;
        int leftLenEnd = data.indexOf(",", valEnd);
        int leftLen = leftLenEnd > 0 ? Integer.parseInt(data.substring(valEnd, leftLenEnd)) : 0;
        int leftStart = leftLen == 0 ? valEnd + 2 : leftLenEnd + 1;

        if (leftLen > 0) {
            node.left = deserialize(data, leftStart, leftStart + leftLen);
        }

        int rightStart = leftLen == 0 ? leftStart : leftStart + leftLen + 1;
        if (start < rightStart && rightStart < end) {
            node.right = deserialize(data, rightStart, end);
        }

        return node;
    }
}
