class Solution:
    """
    https://leetcode.cn/problems/percentage-of-letter-in-string/description/
    """

    def percentageLetter(self, s: str, letter: str) -> int:
        letterCount = 0
        for char in s:
            if char == letter:
                letterCount += 1
        return int(letterCount * 100 / len(s))
