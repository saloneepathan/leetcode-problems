#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<int>> shiftGrid(vector<vector<int>>& grid, int k) {
        int m = grid.size();
        int n = grid[0].size();
        int total = m * n;

        k %= total;

        vector<vector<int>> ans(m, vector<int>(n));

        for (int idx = 0; idx < total; idx++) {
            int newIdx = (idx + k) % total;

            ans[newIdx / n][newIdx % n] =
                grid[idx / n][idx % n];
        }

        return ans;
    }
};

int main() {
    vector<vector<int>> grid = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };

    int k = 1;

    Solution obj;
    vector<vector<int>> result = obj.shiftGrid(grid, k);

    cout << "Shifted Grid:" << endl;
    for (const auto& row : result) {
        for (int num : row) {
            cout << num << " ";
        }
        cout << endl;
    }

    return 0;
}