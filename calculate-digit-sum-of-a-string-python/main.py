def digitSum(s: str, k: int) -> str:
    while len(s) > k:
        new_s = ""

        for i in range(0, len(s), k):
            group = s[i:i + k]

            digit_sum = 0
            for digit in group:
                digit_sum += int(digit)

            new_s += str(digit_sum)

        s = new_s

    return s


if __name__ == "__main__":
    s = "11111222223"
    k = 3

    result = digitSum(s, k)

    print(result)