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
    vector<vector<int>> findDifference(vector<int> &nums1, vector<int> &nums2)
    {
        vector<vector<int>> result;
        result.push_back({});
        result.push_back({});
        unordered_set<int> s1(nums1.begin(), nums1.end());
        unordered_set<int> s2(nums2.begin(), nums2.end());
        for (auto itr = s1.begin(); itr != s1.end(); itr++)
        {
            if (s2.find(*itr) == s2.end())
            {
                result[0].push_back(*itr);
            }
        }
        for (auto itr = s2.begin(); itr != s2.end(); itr++)
        {
            if (s1.find(*itr) == s1.end())
            {
                result[1].push_back(*itr);
            }
        }

        return result;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}