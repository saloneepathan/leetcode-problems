def groupAnagrams(strs: list[str]) -> list[list[str]]:
    group = {}

    for word in strs:
        key = "".join(sorted(word))

        if key not in group:
            group[key] = []
        
        group[key].append(word)

    return list(group.values())

print(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))  # Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]