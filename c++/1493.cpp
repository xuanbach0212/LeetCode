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
    int longestSubarray(vector<int> &nums)
    {
        int leftIdx = 0;
        int rightIdx = 0;
        int maxLen = 0;
        int deleteCount = 0;
        // [1,1,0,1]
        while (rightIdx < nums.size())
        {
            if (nums[rightIdx] == 0)
            {
                deleteCount++;
            }
            while (deleteCount > 1)
            {
                if (nums[leftIdx] == 0)
                {
                    deleteCount--;
                }
                leftIdx++;
            }
            int currLen = rightIdx - leftIdx;
            maxLen = (maxLen > currLen) ? maxLen : currLen;
            rightIdx++;
        }
        return maxLen;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}