#include <iostream>
#include <set>
#include <string>
#include <vector>
using namespace std;

class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        // loop through nums
        for (int i = 0; i < nums.size(); i++)
        {
            for (int j = nums.size() - 1; j > i; j--)
            {
                // left = el, right = find item with value = target - left
                if (nums[j] + nums[i] == target)
                {
                    // if match return [left, right]
                    return {i, j};
                };
            };
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