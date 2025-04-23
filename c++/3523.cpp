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
    int maximumPossibleSize(vector<int> &nums)
    {
        // [10,43,43,17,36,44,35,28]
        int removed = 0;
        int last = -1;
        for (int num : nums)
        {
            if (num >= last)
            {
                last = num;
            }
            else
            {
                removed++;
            }
        }

        return nums.size() - removed;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}