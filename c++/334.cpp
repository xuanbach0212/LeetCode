#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>
#include <unordered_set>
#include <stack>
#include <unordered_map>

using namespace std;

class Solution
{
public:
    bool increasingTriplet(vector<int> &nums)
    {
        // Input: nums = [1,2,3,4,5]
        // Output: true
        // Explanation: Any triplet where i < j < k is valid.

        // [20,100,10,12,5,100,4,3,13]
        // [5,4,3,2,1]
        // [2,1,5,0,4,6]
        // [1,5,0,4,1,3]
        // [4,5,2147483647,1,2]
        // [1,0,4,2,1,0,-1,-3]
        // [1,5,0,4,1,3]
        if (nums.size() < 3)
            return false;
        int first = nums[0];
        int second;
        bool hasSecond = false;
        for (int i = 1; i < nums.size(); i++)
        {
            int num = nums[i];
            if (num <= first)
            {
                first = num;
            }
            else if (!hasSecond)
            {
                second = num;
                hasSecond = true;
            }
            else if (num <= second)
            {
                second = num;
            }
            else
            {
                return true;
            }
        }
        return false;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}