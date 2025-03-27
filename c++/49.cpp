#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    vector<vector<string>> groupAnagrams(vector<string> &strs)
    {
        vector<vector<string>> ans;
        map<string, vector<string>> m;

        // make key like "meet" -> 1m2e3t and value is match this key
        for (int i = 0; i < strs.size(); i++)
        {
            map<char, int> keyCharMap;
            // loop strs to separate char
            for (char c : strs[i])
            {
                keyCharMap[c]++;
            }
            string ansKey = "";
            for (auto [key, value] : keyCharMap)
            {
                ansKey += to_string(value) + key;
            };
            m[ansKey].push_back(strs[i]);
        };

        for (auto [key, value] : m)
        {
            ans.push_back(value);
        };

        return ans;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}