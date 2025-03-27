#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    vector<int> getConcatenation(vector<int> &nums)
    {
        // declare vector ans
        int len = nums.size();
        vector<int> ans(2 * len, 0);
        // loop through nums
        for (int i = 0; i < len; i++)
        {
            // add ans[i] = nums[i] and ans[i+n]=nums[i] at the same time
            ans.at(i) = nums[i];
            ans.at(i + len) = nums[i];
        };
        return ans;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}