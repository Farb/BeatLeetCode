/**
 * https://leetcode.cn/problems/percentage-of-letter-in-string/description/?envType=daily-question&envId=2025-03-31
 */
public class P2278 {
    public int percentageLetter(String s, char letter) {
        int letterCount = 0;
        for (char c : s.toCharArray()) {
            if (c == letter) {
                letterCount++;
            }
        }
        return letterCount * 100 / s.length();
    }
}
