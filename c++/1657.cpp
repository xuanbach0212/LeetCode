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
    bool closeStrings(string word1, string word2)
    {
        unordered_map<char, int> m1;
        unordered_map<char, int> m2;

        for (char c : word1)
        {
            m1[c]++;
        }
        for (char c : word2)
        {
            if (!m1.count(c))
            {
                return false;
            }
            m2[c]++;
        }
        if (m1.size() != m2.size())
            return false;

        multiset<int> s1, s2;
        for (auto &p : m1)
            s1.insert(p.second);
        for (auto &p : m2)
            s2.insert(p.second);

        return s1 == s2;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}