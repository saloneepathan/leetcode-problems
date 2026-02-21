#include <iostream>
using namespace std;

class Solution {
public:
    
    bool isPrime(int n)
    {
        if(n < 2) return false;
        
        for(int i = 2; i * i <= n; i++)
        {
            if(n % i == 0)
                return false;
        }
        
        return true;
    }
    
    int countPrimeSetBits(int left, int right) {
        
        int count = 0;
        
        for(int i = left; i <= right; i++)
        {
            int setBits = __builtin_popcount(i);
            
            if(isPrime(setBits))
                count++;
        }
        
        return count;
    }
};

int main()
{
    Solution obj;
    
    int left, right;
    
    cout << "Enter left: ";
    cin >> left;
    
    cout << "Enter right: ";
    cin >> right;
    
    int result = obj.countPrimeSetBits(left, right);
    
    cout << "Output: " << result << endl;
    
    return 0;
}