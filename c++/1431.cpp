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
    vector<bool> kidsWithCandies(vector<int> &candies, int extraCandies)
    {
        int maxCandies = 0;
        vector<bool> results;
        for (int i = 0; i < candies.size(); i++)
        {
            maxCandies = candies[i] > maxCandies ? candies[i] : maxCandies;
        }
        for (int i = 0; i < candies.size(); i++)
        {
            if (candies[i] + extraCandies >= maxCandies)
            {
                results.push_back(true);
            }
            else
            {
                results.push_back(false);
            }
        }
        return results;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}