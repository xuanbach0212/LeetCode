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
    bool containsNearbyDuplicate(vector<int> &nums, int k)
    {
        unordered_set<int> window;
        for (int i = 0; i < nums.size(); i++)
        {
            if (window.count(nums[i]))
            {
                return true;
            }
            window.insert(nums[i]);
            if (window.size() > k)
            {
                window.erase(nums[i - k]);
            }
        }
        return false;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};
    bool result = solution.containsNearbyDuplicate(nums, 3);

    cout << "result" << "\n";
    cout << result;

    return 0;
}
