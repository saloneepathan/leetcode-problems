#include <iostream>
using namespace std;

class Solution {
public:
    int binaryGap(int n) {
        
        int last = -1;
        int maxDist = 0;
        int pos = 0;
        
        while (n > 0) {
            
            if (n & 1) {
                
                if (last != -1) {
                    maxDist = max(maxDist, pos - last);
                }
                
                last = pos;
            }
            
            pos++;
            n = n >> 1;
        }
        
        return maxDist;
    }
};

int main() {
    
    Solution sol;
    
    int n;
    cout << "Enter number: ";
    cin >> n;
    
    int result = sol.binaryGap(n);
    
    cout << "Binary Gap: " << result << endl;
    
    return 0;
}