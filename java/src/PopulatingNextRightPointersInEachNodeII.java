/*
 * 使用链表作队列来层序遍历，并使用一个int来指示每轮需要出队的结点数（即每层的结点数），
 * 每结点出队后以其填充上一个出队的结点的next指针
 */
public class PopulatingNextRightPointersInEachNodeII {
    public static void main(String[] args) {
        Node root = new Node(
                1,
                new Node(2,
                        new Node(4, null, null, null),
                        new Node(5, null, null, null),
                        null),
                new Node(3,
                        null,
                        new Node(7, null, null, null),
                        null),
                null
        );

        new PopulatingNextRightPointersInEachNodeII().connect(root);

        Node sibling = null;
        while (null != root) {
            sibling = root;
            while (null != sibling) {
                System.out.println(sibling.val);
                sibling = sibling.next;
            }
            root = root.left;
        }
    }

    public Node connect(Node root) {
        Node level_start = root;
        while (level_start != null) {
            Node cur = level_start;
            while (cur != null) {
                if (cur.left != null) {
                    cur.left.next = cur.right == null ? siblingFirstChild(cur.next) : cur.right;
                }
                if (cur.right != null) {
                    cur.right.next = siblingFirstChild(cur.next);
                }
                cur = cur.next;
            }
            level_start = siblingFirstChild(level_start);
        }
        return root;
    }

    private Node siblingFirstChild(Node node) {
        while (node != null) {
            if (node.left != null) return node.left;
            if (node.right != null) return node.right;
            node = node.next;
        }
        return null;
    }
}
