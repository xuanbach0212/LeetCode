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
    bool isSubsequence(string s, string t)
    {
        int i = 0;
        int j = 0;
        while (i < s.size(), j < t.size())
        {
            if (s[i] == t[j])
            {
                i++;
            }
            j++;
        }
        return i == s.size();
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}