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
    int equalPairs(vector<vector<int>> &grid)
    {
        // [[11,1],[1,11]]
        int n = grid.size();
        unordered_map<string, int> m;
        for (int i = 0; i < n; i++)
        {
            string k = "";
            for (int v : grid[i])
            {
                k += to_string(v) + "_";
            }
            m[k]++;
        }

        int ans = 0;

        for (int i = 0; i < n; i++)
        {
            vector<int> columns;
            for (int j = 0; j < n; j++)
            {
                columns.push_back(grid[j][i]);
            }
            string k = "";
            for (int v : columns)
            {
                k += to_string(v) + "_";
            }
            if (m.count(k))
            {
                ans += m[k];
            }
        }
        return ans;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}