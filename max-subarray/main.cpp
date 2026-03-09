#include <iostream>
#include <vector>
using namespace std;

int maxSubarray(vector<int>& nums){
    
    int maxSum = nums[0];
    int currSum = nums[0];
    
    for (int i = 0; i<nums.size(); i++){
        currSum = max(nums[i], currSum + nums[i]);
        maxSum = max(currSum, maxSum);
    }
    
    return maxSum;
}

int main(){
    
    vector<int> nums = {2, 5, 3, -4, -1, 2, -1, 5, -9, 10, -3, -5};
    
    cout << "Maximum subarray sum is " << maxSubarray(nums);
    
    return 0;
}