#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    vector<int> productExceptSelf(vector<int> &nums)
    {
        vector<int> v(nums.size(), 1);
        long long prefix = 1, suffix = 1;
        for (int i = 0; i < nums.size(); i++)
        {
            v[i] *= prefix;
            prefix *= nums[i];
        }
        for (int i = nums.size() - 1; i >= 0; i--)
        {
            v[i] *= suffix;
            suffix *= nums[i];
        }
        return v;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}