#include <iostream>
#include <vector>
#include <stack>
#include <algorithm>
#include <string>

using namespace std;

class Solution {
public:
    vector<int> survivedRobotsHealths(vector<int>& positions,
                                      vector<int>& healths,
                                      string directions) {
        int n = positions.size();
        vector<int> indices(n), result;
        stack<int> st;

        for (int i = 0; i < n; i++) {
            indices[i] = i;
        }

        sort(indices.begin(), indices.end(),
             [&](int a, int b) {
                 return positions[a] < positions[b];
             });

        for (int currentIndex : indices) {
            if (directions[currentIndex] == 'R') {
                st.push(currentIndex);
            } else {
                while (!st.empty() && healths[currentIndex] > 0) {
                    int topIndex = st.top();
                    st.pop();

                    if (healths[topIndex] > healths[currentIndex]) {
                        healths[topIndex] -= 1;
                        healths[currentIndex] = 0;
                        st.push(topIndex);
                    } 
                    else if (healths[topIndex] < healths[currentIndex]) {
                        healths[currentIndex] -= 1;
                        healths[topIndex] = 0;
                    } 
                    else {
                        healths[currentIndex] = 0;
                        healths[topIndex] = 0;
                    }
                }
            }
        }

        for (int i = 0; i < n; i++) {
            if (healths[i] > 0) {
                result.push_back(healths[i]);
            }
        }

        return result;
    }
};

int main() {
    Solution sol;

    vector<int> positions = {5, 4, 3, 2, 1};
    vector<int> healths = {2, 17, 9, 15, 10};
    string directions = "RRRRR";

    vector<int> result = sol.survivedRobotsHealths(positions, healths, directions);

    cout << "Surviving robots healths: ";
    for (int x : result) {
        cout << x << " ";
    }
    cout << endl;

    return 0;
}