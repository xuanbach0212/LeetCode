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
    int minMoves(int target, int maxDoubles)
    {
        int stepCount = 0;
        if (maxDoubles == 0)
        {
            return target - 1;
        }
        while (target > 1)
        {
            if (maxDoubles == 0)
            {
                return stepCount + target - 1;
            }

            if (maxDoubles > 0 && target % 2 == 0)
            {
                target = target / 2;
                maxDoubles -= 1;
                stepCount += 1;
                continue;
            }
            target -= 1;
            stepCount += 1;
        }
        return stepCount;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}