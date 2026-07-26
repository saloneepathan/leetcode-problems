#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    const int MOD = 1000000007;

    int numberOfStableArrays(int zero, int one, int limit) {

        vector<vector<vector<long long>>> dp(
            zero + 1,
            vector<vector<long long>>(one + 1, vector<long long>(2, 0))
        );

        // Start with zeros
        for(int k = 1; k <= min(limit, zero); k++) {
            dp[k][0][0] = 1;
        }

        // Start with ones
        for(int k = 1; k <= min(limit, one); k++) {
            dp[0][k][1] = 1;
        }

        for(int z = 0; z <= zero; z++) {
            for(int o = 0; o <= one; o++) {

                if(dp[z][o][0] > 0) {
                    for(int k = 1; k <= limit && o + k <= one; k++) {
                        dp[z][o + k][1] = (dp[z][o + k][1] + dp[z][o][0]) % MOD;
                    }
                }

                if(dp[z][o][1] > 0) {
                    for(int k = 1; k <= limit && z + k <= zero; k++) {
                        dp[z + k][o][0] = (dp[z + k][o][0] + dp[z][o][1]) % MOD;
                    }
                }

            }
        }

        return (dp[zero][one][0] + dp[zero][one][1]) % MOD;
    }
};

int main() {
    Solution sol;

    int zero, one, limit;

    cout << "Enter number of zeros: ";
    cin >> zero;

    cout << "Enter number of ones: ";
    cin >> one;

    cout << "Enter limit: ";
    cin >> limit;

    int result = sol.numberOfStableArrays(zero, one, limit);

    cout << "Number of stable arrays: " << result << endl;

    return 0;
}