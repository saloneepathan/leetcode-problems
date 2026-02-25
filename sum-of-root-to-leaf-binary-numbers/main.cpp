#include <iostream>
#include <vector>
#include <queue>
using namespace std;

/**
 * Definition for a binary tree node.
 */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode(int x) {
        val = x;
        left = NULL;
        right = NULL;
    }
};

class Solution {
public:

    int dfs(TreeNode* node, int current)
    {
        if(node == NULL)
            return 0;

        current = (current << 1) | node->val;

        // if leaf
        if(node->left == NULL && node->right == NULL)
            return current;

        return dfs(node->left, current) + dfs(node->right, current);
    }

    int sumRootToLeaf(TreeNode* root) {
        return dfs(root, 0);
    }
};


// Function to build tree from level order array (-1 means NULL)
TreeNode* buildTree(vector<int>& arr)
{
    if(arr.size() == 0 || arr[0] == -1)
        return NULL;

    TreeNode* root = new TreeNode(arr[0]);
    queue<TreeNode*> q;
    q.push(root);

    int i = 1;

    while(!q.empty() && i < arr.size())
    {
        TreeNode* current = q.front();
        q.pop();

        // left child
        if(arr[i] != -1)
        {
            current->left = new TreeNode(arr[i]);
            q.push(current->left);
        }
        i++;

        if(i >= arr.size())
            break;

        // right child
        if(arr[i] != -1)
        {
            current->right = new TreeNode(arr[i]);
            q.push(current->right);
        }
        i++;
    }

    return root;
}


int main()
{
    // Example: [1,0,1,0,1,0,1]

    vector<int> arr = {1,0,1,0,1,0,1};

    TreeNode* root = buildTree(arr);

    Solution obj;

    int result = obj.sumRootToLeaf(root);

    cout << "Output: " << result << endl;

    return 0;
}