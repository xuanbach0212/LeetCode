#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    int pivotIndex(vector<int> &nums)
    {
        vector<int> prefixSum;
        int sum = 0;
        for (int num : nums)
        {
            sum += num;
            prefixSum.push_back(sum);
        }
        for (int i = 0; i < nums.size(); i++)
        {
            int left, right;
            left = (i != 0) ? prefixSum[i - 1] : 0;

            if (i != nums.size() - 1)
            {
                int minusNum = prefixSum[nums.size() - 1];
                int beMinusNum = prefixSum[i];
                right = minusNum - beMinusNum;
            }
            else
            {
                right = 0;
            }

            if (left == right)
            {
                return i;
            }
        }
        return -1;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 7, 3, 6, 5, 6};
    int result = solution.pivotIndex(nums);
    cout << "Result: " << result << endl;

    return 0;
}