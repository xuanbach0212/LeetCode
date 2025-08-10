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
    double findMaxAverage(vector<int> &nums, int k)
    {
        int i = 0;
        int j = 0;
        int maxOne = nums[0];
        int temp = nums[0];
        while (j < nums.size() || i < nums.size())
        {
            if (j - i < k)
            {
                maxOne += nums[j];
                temp += nums[j];
            }
            else if (j - i == k)
            {
                temp -= nums[i];
                temp += nums[j];
                maxOne = (maxOne > temp) ? maxOne : temp;
                i++;
            }
            j++;
        }
        double(maxOne) / k;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}