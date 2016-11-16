#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int findPeakElement(vector<int>& nums) {
        return iterator(nums, 0, nums.size() - 1);
    }
private:
    int iterator(vector<int>& vec, int left, int right) {
        if (right - left <= 1) {
            return vec[right] > vec[left] ? right : left;
        }

        int middle = (left+right)/2;
        if (vec[middle - 1] > vec[middle]) {
            return iterator(vec, left, middle);
        }
        if (vec[middle + 1] > vec[middle]) {
            return iterator(vec, middle, right);
        }
        return middle;
    }
};

int main(void) {
    Solution* s = new Solution();
    vector<int> vec;
    vec.push_back(17);
    vec.push_back(18);
    vec.push_back(19);
    vec.push_back(20);
    vec.push_back(21);
    vec.push_back(22);
    vec.push_back(11);
    vec.push_back(12);
    vec.push_back(13);
    vec.push_back(14);
    vec.push_back(15);
    vec.push_back(14);

    cout << s->findPeakElement(vec) << "\n";
    return 0;
}
