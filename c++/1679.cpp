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
    int maxOperations(vector<int> &nums, int k)
    {
        // [3,1,3,4,3]
        // [1,2,3,4]
        unordered_map<int, int> m;
        int count = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            if (m.count(k - nums[i]) && m[k - nums[i]] > 0)
            {
                count++;
                m[k - nums[i]]--;
            }
            else
            {
                m[nums[i]]++;
            }
        }
        return count;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}