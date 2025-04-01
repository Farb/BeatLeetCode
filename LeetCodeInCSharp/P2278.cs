namespace LeetCodeInCSharp;

/// <summary>
/// https://leetcode.cn/problems/percentage-of-letter-in-string/description/?envType=daily-question&envId=2025-03-31
/// </summary>
internal class P2278
{
    public int PercentageLetter(string s, char letter)
    {
        int letterCount = 0;
        foreach (char c in s)
        {
            if (char.Equals(c, letter))
            {
                letterCount++;
            }
        }
        return letterCount * 100 / s.Length;
    }
}
