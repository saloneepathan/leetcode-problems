#include <iostream>
#include <unordered_set>
using namespace std;

int longestLengthOfSubstring(string str){
        unordered_set<char> set;
        int left = 0;
        int maxLength = 0;
        
        for (int right = 0; right < str.length(); right++){
            while (set.find(str[right]) != set.end()){
                set.erase(str[left]);
                left++;
            }
            set.insert(str[right]);
            maxLength = max(maxLength, right-left+1);
        }
        return maxLength;
}

int main(){
    string str = "abcabcbb";
    int result = longestLengthOfSubstring(str);
    
    cout << "Longest Length of the substring is " << result << endl;
    return 0;
}