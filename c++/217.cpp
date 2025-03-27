#include <iostream>
using std::vector;
#include <set>

class Solution
{
public:
    bool containsDuplicate(vector<int> &nums)
    {
        std::set<int> s(nums.begin(), nums.end());
        if (s.size() != nums.size())
        {
            return true;
        }
        return false;
    }
};
