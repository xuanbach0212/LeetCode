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
    int climbStairs(int n)
    {
        if (n == 0 || n == 1)
        {
            return 1;
        }
        int prev = 1, curr = 1;
        for (int i = 2; i <= n; i++)
        {
            int temp = curr;
            curr = prev + temp;
            prev = temp;
        }
        return curr;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};
    // bool result = solution.containsNearbyDuplicate(nums, 3);

    cout << "result" << "\n";
    // cout << result;

    return 0;
}
