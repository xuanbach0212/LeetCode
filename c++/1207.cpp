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
    bool uniqueOccurrences(vector<int> &arr)
    {
        unordered_map<int, int> m;
        unordered_set<int> s;
        for (int v : arr)
        {
            m[v]++;
        }
        for (auto it = m.begin(); it != m.end(); it++)
        {
            s.insert(it->second);
        }

        return s.size() == m.size();
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}