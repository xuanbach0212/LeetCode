#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>
#include <unordered_set>

using namespace std;

class Solution
{
public:
    int longestConsecutive(vector<int> &nums)
    {
        // remove duplicate
        unordered_set<int> s(nums.begin(), nums.end());

        int biggestStreak = 0;
        for (int el : s)
        {
            // if is start of sequence
            if (s.find(el - 1) == s.end())
            {
                int streak = 0;
                while (s.find(el + streak) != s.end())
                {
                    streak += 1;
                }
                biggestStreak = max(biggestStreak, streak);
            }
        }
        return biggestStreak;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}