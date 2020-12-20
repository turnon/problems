import java.util.LinkedList;

/*
 * 使用链表作队列来层序遍历，并使用一个int来指示每轮需要出队的结点数（即每层的结点数），
 * 每结点出队后以其填充上一个出队的结点的next指针
 */
public class PopulatingNextRightPointersInEachNode {
    public static void main(String[] args) {
        Node root = new Node(
                1,
                new Node(2,
                        new Node(4, null, null, null),
                        new Node(5, null, null, null),
                        null),
                new Node(3,
                        new Node(6, null, null, null),
                        new Node(7, null, null, null),
                        null),
                null
        );

        new PopulatingNextRightPointersInEachNode().connect(root);

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
        if (null == root) {
            return null;
        }

        LinkedList<Node> q = new LinkedList<Node>();
        int nodesInLevel = 1;
        q.add(root);
        while (!q.isEmpty()) {
            deqAndEnq(q, nodesInLevel);
            nodesInLevel = nodesInLevel << 1;
        }

        return root;
    }

    private void deqAndEnq(LinkedList<Node> q, int num) {
        Node last = null;
        while (num > 0) {
            Node head = q.removeFirst();
            if (null != last) {
                last.next = head;
            }
            last = head;
            if (null != last.left) {
                q.add(last.left);
                q.add(last.right);
            }
            num--;
        }
    }

//    /*
//     * 论坛的解法没有使用辅助的数据结构，号称空间复杂度O(1)，但结果好像没什么提高…
//     */
//    public Node connect(Node root) {
//        Node level_start = root;
//        while (level_start != null) {
//            Node cur = level_start;
//            while (cur != null) {
//                if (cur.left != null) cur.left.next = cur.right;
//                if (cur.right != null && cur.next != null) cur.right.next = cur.next.left;
//
//                cur = cur.next;
//            }
//            level_start = level_start.left;
//        }
//        return root;
//    }
}
