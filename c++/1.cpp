#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        // make a map
        map<int, int> m;
        // loop nums if right = target - left
        for (int i = 0; i < nums.size(); i++)
        {
            int right = target - nums[i];
            // find if right in map -> if find return
            if (m.find(right) != m.end())
            {
                return {i, m[right]};
            };
            // if not find insert left for next el
            m.insert({nums[i], i});
        };
        return {};
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}