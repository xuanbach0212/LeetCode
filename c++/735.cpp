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
    vector<int> asteroidCollision(vector<int> &asteroids)
    {
        // Input: asteroids = [5,10,-5]
        // Output: [5,10]
        // Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
        vector<int> stack;
        int n = asteroids.size();
        int left = 0;
        while (left <= n - 1)
        {
            int curr = asteroids[left];
            if (stack.size() == 0)
            {
                stack.push_back(curr);
                left++;
                continue;
            }
            int lastInStack = stack[stack.size() - 1];
            if (curr < 0 && lastInStack > 0)
            {
                if (lastInStack + curr > 0)
                {
                    left++;
                }
                else if (lastInStack + curr < 0)
                {
                    stack.pop_back();
                }
                else
                {
                    stack.pop_back();
                    left++;
                }
                continue;
            }
            stack.push_back(curr);
            left++;
            continue;
        }
        return stack;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}