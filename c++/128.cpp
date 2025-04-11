#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    int longestConsecutive(vector<int> &nums)
    {
        // remove duplicate and sort
        set<int> s(nums.begin(), nums.end());

        int biggestStreak = 0;
        int streak = 1;
        int currentValue = *s.begin();
        for (int el : s)
        {
            streak = (el - currentValue == 1) ? streak + 1 : 1;
            currentValue = el;
            biggestStreak = (biggestStreak > streak) ? biggestStreak : streak;
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