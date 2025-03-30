#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    vector<int> topKFrequent(vector<int> &nums, int k)
    {
        map<int, int> m;
        map<int, vector<int>> fm;
        vector<int> v;

        // key is number, value is frequent
        for (int num : nums)
        {
            m[num]++;
        }
        // key is frequent, value is vector of key
        for (auto [key, value] : m)
        {
            fm[value].push_back(key);
        };
        for (auto it = fm.rbegin(); it != fm.rend(); it++)
        {
            for (int el : it->second)
            {
                if (v.size() >= k)
                {
                    return v;
                }
                v.push_back(el);
            };
            if (v.size() >= k)
            {
                return v;
            }
        };
        return v;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}