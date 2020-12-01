public class ConvertSortedArrayToBinarySearchTree {
    public static void main(String[] args) {
        int[] nums = new int[]{-10, -3, 0, 5, 9};
        TreeNode node = new ConvertSortedArrayToBinarySearchTree().sortedArrayToBST(nums);
        System.out.println(node.inorder());
    }

    public TreeNode sortedArrayToBST(int[] nums) {
        if (nums.length == 0) return null;
        return sortedArrayToBST(nums, 0, nums.length - 1);
    }

    public TreeNode sortedArrayToBST(int[] nums, int start, int end) {
        int distance = end - start;
        if (distance == 0) {
            return new TreeNode(nums[start], null, null);
        } else if (distance == 1) {
            return new TreeNode(nums[end], new TreeNode(nums[start], null, null), null);
        }

        int mid = (end + start) / 2;

        return new TreeNode(
                nums[mid],
                sortedArrayToBST(nums, start, mid - 1),
                sortedArrayToBST(nums, mid + 1, end)
        );
    }
}
