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
    void moveZeroes(vector<int> &nums)
    {
        int left = 0;
        int right = 0;
        while (right < nums.size())
        {
            // [1,0,0,3,12]
            // [1,0,1]
            if (nums[right] != 0)
            {
                if (nums[left] == 0)
                {
                    int temp = nums[right];
                    nums[right] = nums[left];
                    nums[left] = temp;
                }
                left++;
            }
            right++;
        }
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}