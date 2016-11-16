#include <iostream>
#include <vector>
using namespace std;

typedef struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
} TreeNode;

class Solution {
public:
    Solution() {

    }
    vector<string> binaryTreePaths(TreeNode* root) {
        string s = "";
        this->_iterator(root, s);
        return this->_vector;
    }

private:
    void _iterator(TreeNode* root, string s) {
        if (root == NULL) {
            return;
        }
        s += this->_itoa(root->val);
        if (root->left == NULL && root->right == NULL) {
            this->_vector.push_back(s);
        }
        s += "->";
        if (root->left != NULL) {
            this->_iterator(root->left, s);
        }
        if (root->right != NULL) {
            this->_iterator(root->right, s);
        }
    }
    string _itoa(int i) {
        string result = "";
        bool is_negative = false;
        if (i < 0) {
            is_negative = true;
            i = -i;
        }
        while (i > 0) {
            result = (char)((i%10) + '0') + result;
            i = i / 10;
        }
        if (is_negative) {
            result = "-" + result;
        }
        return result;
    }
    vector<string> _vector;
};

int main(void) {
    Solution* s = new Solution();
    TreeNode* root = new TreeNode();
    root->val = 1;
    root->left = new TreeNode();
    root->right = new TreeNode();
    root->left->val = 2;
    root->right->val = 3;
    root->left->right = new TreeNode();
    root->left->right->val = 5;
    root->left->right->left = new TreeNode();
    root->left->right->left->val = 10;
    root->left->right->left->left = new TreeNode();
    root->left->right->left->left->val = 12;
    vector<string> v = s->binaryTreePaths(root);
    for (vector<string>::iterator it = v.begin();
         it != v.end(); it++) {
        cout << *it << "\n";
    }
    return 0;
}
