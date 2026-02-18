#include <iostream>
using namespace std;

class Solution {
public:
    bool hasAlternatingBits(int n) {

        unsigned int x = n ^ (n >> 1);

        return (x & (x + 1)) == 0;
    }
};

int main() {

    Solution sol;
    int n;
    cin >> n;

    cout << (sol.hasAlternatingBits(n) ? "true" : "false");

    return 0;
}
