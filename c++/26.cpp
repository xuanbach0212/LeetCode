#include <iostream>
#include <set>
#include <string>
#include <vector>
using namespace std;

class Solution
{
public:
    int removeDuplicates(vector<int> &nums)
    {
        set<int> s;
        for (int i = 0; i < nums.size(); i++)
        {
            s.insert(nums[i]);
        };
        nums.clear();
        for (auto sc : s)
        {
            nums.push_back(sc);
        };
        return s.size();

        // nums.erase(unique(nums.begin(), nums.end()), nums.end());
        // return nums.size();
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};
    
    cout << "Original array: ";
    for (int num : nums) {
        cout << num << " ";
    }
    cout << endl;
    
    int newLength = solution.removeDuplicates(nums);
    
    cout << "After removing duplicates: ";
    for (int num : nums) {
        cout << num << " ";
    }
    cout << endl;
    cout << "New length: " << newLength << endl;

    return 0;
}