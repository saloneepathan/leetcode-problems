#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int numberOfSubmatrices(vector<vector<char>>& grid) {
        int rows = grid.size(), cols = grid[0].size();
        vector<int> sumX(cols, 0);
        vector<int> sumY(cols, 0);
        int res = 0;

        for (int i = 0; i < rows; i++) {
            int rx = 0, ry = 0;

            for (int j = 0; j < cols; j++) {
                if (grid[i][j] == 'X') rx++;
                else if (grid[i][j] == 'Y') ry++;

                sumX[j] += rx;
                sumY[j] += ry;

                if (sumX[j] > 0 && sumX[j] == sumY[j]) {
                    res++;
                }
            }
        }

        return res;
    }
};

int main() {
    Solution sol;

    vector<vector<char>> grid1 = {
        {'X','Y','.'},
        {'Y','.','.'}
    };
    cout << sol.numberOfSubmatrices(grid1) << endl; // 3

    vector<vector<char>> grid2 = {
        {'X','X'},
        {'X','Y'}
    };
    cout << sol.numberOfSubmatrices(grid2) << endl; // 0

    vector<vector<char>> grid3 = {
        {'.','.'},
        {'.','.'}
    };
    cout << sol.numberOfSubmatrices(grid3) << endl; // 0

    return 0;
}