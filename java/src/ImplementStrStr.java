public class ImplementStrStr {
    public static void main(String[] args) {
        ImplementStrStr iss = new ImplementStrStr();
        System.out.println(iss.strStr("hello", "ll"));
        System.out.println(iss.strStr("aaaaa", "bba"));
        System.out.println(iss.strStr("aaaaa", "aaaaaa"));
        System.out.println(iss.strStr("", ""));
    }

    public int strStr(String haystack, String needle) {
        if ("".equals(haystack) || "".equals(needle)) return 0;

        int last_head_pos = haystack.length() - needle.length();
        if (last_head_pos < 0) return -1;

        char head = needle.charAt(0);
        for (int i = 0; i < last_head_pos; i++) {
            if (haystack.charAt(i) == head && matchRest(haystack, i, needle)) {
                return i;
            }
        }

        return -1;
    }

    private boolean matchRest(String haystack, int start, String needle) {
        for (int i = 1; i < needle.length(); i++) {
            if (haystack.charAt(start + i) != needle.charAt(i)) {
                return false;
            }
        }
        return true;
    }
}
