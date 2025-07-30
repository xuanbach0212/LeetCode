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
    int numTilePossibilities(string tiles)
    {
        unordered_map<char, int> m;
        for (char c : tiles)
        {
            m[c] += 1;
        }
        return backtracking(m);
    }

    int backtracking(unordered_map<char, int> m)
    {
        int subRes = 0;
        for (auto [c, count] : m)
        {
            if (count > 0)
            {
                m[c] -= 1;
                subRes += 1;
                subRes += backtracking(m);
                m[c] += 1;
            }
        }
        return subRes;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}