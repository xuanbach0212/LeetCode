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
    int longestOnes(vector<int> &nums, int k)
    {
        // Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
        // Output: 6
        // Explanation: [1,1,1,0,0,1,1,1,1,1,1]
        int leftIdx = 0;
        int rightIdx = 0;
        int zeroCount = 0;
        int maxLen = 0;
        while (rightIdx < nums.size())
        {
            if (nums[rightIdx] == 0)
            {
                zeroCount++;
            }
            while (zeroCount > k)
            {
                if (nums[leftIdx] == 0)
                {
                    zeroCount--;
                }
                leftIdx++;
            }
            int currLen = rightIdx - leftIdx + 1;
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