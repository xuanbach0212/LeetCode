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
    int maxVowels(string s, int k)
    {
        unordered_set<char> vowels = {'a', 'i', 'u', 'e', 'o'};
        // "weallloveyou"
        // "abciiidef"
        // 1 1 1 2 3 4 4 5 5
        vector<int> prefixCount;
        int count = 0;
        for (int i = 0; i < s.size(); i++)
        {
            auto it = vowels.find(tolower(s[i]));
            if (it != vowels.end())
            {
                count += 1;
            }
            prefixCount.push_back(count);
        }

        // a e i o u with k = 2
        // 1 2 3 4 5
        int maxCount = 0;
        for (int i = 0; i < s.size(); i++)
        {
            int count;
            int rightIdx = i + k - 1;
            int leftIdx = i - 1;
            if (rightIdx > s.size() - 1 || maxCount == k)
            {
                return maxCount;
            }
            int left = (leftIdx < 0) ? 0 : prefixCount[leftIdx];
            count = prefixCount[rightIdx] - left;
            maxCount = (maxCount > count) ? maxCount : count;
        }

        return maxCount;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}