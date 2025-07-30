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
    string mergeAlternately(string word1, string word2)
    {
        string merged;
        int i = 0;
        while (i < max(word2.size(), word1.size()))
        {
            if (i < word1.size())
            {
                merged += word1[i];
            }
            if (i < word2.size())
            {
                merged += word2[i];
            }
            i++;
        }

        return merged;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}