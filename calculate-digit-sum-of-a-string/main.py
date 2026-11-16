class Solution:
    def digitSum(self, s: str, k: int) -> str:
        # Repeat rounds while the string is longer than k
        while len(s) > k:
            parts = []
            for i in range(0, len(s), k):
                # sum the digits of one group (last group may be shorter)
                parts.append(str(sum(int(c) for c in s[i:i + k])))
            s = "".join(parts)
        return s


if __name__ == "__main__":
    sol = Solution()
    tests = [
        ("11111222223", 3, "135"),  # example 1
        ("00000000", 3, "000"),     # example 2
        ("1", 2, "1"),              # edge: already short enough
        ("12", 2, "12"),            # edge: length == k, no round
        ("999", 2, "99"),           # "189" -> "99"
        ("9" * 100, 2, "99"),       # max size
    ]
    for s, k, exp in tests:
        act = sol.digitSum(s, k)
        status = "OK" if act == exp else "FAIL"
        print(f"input: s={s[:20]!r}, k={k} | expected: {exp} | actual: {act} | {status}")
