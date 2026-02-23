#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
public:
    bool hasAllCodes(string s, int k) {
        int n = s.size();

        int total = 1 << k;

        vector<bool> seen(total, false);

        int mask = total - 1;
        int num = 0;
        int count = 0;

        for (int i = 0; i < n; i++) {

            num = ((num << 1) & mask) | (s[i] - '0');

            if (i >= k - 1) {

                if (!seen[num]) {
                    seen[num] = true;
                    count++;

                    if (count == total)
                        return true;
                }
            }
        }

        return false;
    }
};

int main() {

    Solution sol;

    string s;
    int k;

    cout << "Enter binary string: ";
    cin >> s;

    cout << "Enter k: ";
    cin >> k;

    if (sol.hasAllCodes(s, k))
        cout << "true" << endl;
    else
        cout << "false" << endl;

    return 0;
}