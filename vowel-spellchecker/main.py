from typing import List

class Solution:
    def spellchecker(self, wordlist: List[str], queries: List[str]) -> List[str]:
        vowels = set("aeiou")

        # Replace all vowels in a word with '*'
        def devowel(word: str) -> str:
            return ''.join('*' if c in vowels else c for c in word)

        # For exact match lookups
        word_set = set(wordlist)

        # Maps for case-insensitive and vowel-error matches
        case_insensitive_map = {}
        vowel_error_map = {}

        # Build the maps (only store the first word encountered for precedence)
        for word in wordlist:
            lower = word.lower()
            case_insensitive_map.setdefault(lower, word)
            vowel_error_map.setdefault(devowel(lower), word)

        ans = []
        for query in queries:
            if query in word_set:  
                # Rule 1: Exact match
                ans.append(query)
            else:
                lower = query.lower()
                if lower in case_insensitive_map:
                    # Rule 2: Case-insensitive match
                    ans.append(case_insensitive_map[lower])
                else:
                    devoweled = devowel(lower)
                    if devoweled in vowel_error_map:
                        # Rule 3: Vowel error match
                        ans.append(vowel_error_map[devoweled])
                    else:
                        # Rule 4: No match
                        ans.append("")
        return ans


# -----------------------
# Example usage:
# -----------------------
if __name__ == "__main__":
    solution = Solution()

    wordlist = ["KiTe","kite","hare","Hare"]
    queries = ["kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"]
    print(solution.spellchecker(wordlist, queries))
    # Output: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]

    wordlist2 = ["yellow"]
    queries2 = ["YellOw"]
    print(solution.spellchecker(wordlist2, queries2))
    # Output: ["yellow"]
