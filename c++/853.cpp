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
    // Input: target = 100, position = [0,2,4], speed = [4,2,1]
    // Output: 1
    // Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
    // Output: 3
    // Explanation:
    // The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12. The fleet forms at target.
    // The car starting at 0 (speed 1) does not catch up to any other car, so it is a fleet by itself.
    // The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.
    int carFleet(int target, vector<int> &position, vector<int> &speed)
    {
        map<int, int> m;
        vector<float> stack;
        for (int i = 0; i < position.size(); i++)
        {
            m.insert({position[i], speed[i]});
        }

        // [3, 2.5]
        for (auto iter = m.rbegin(); iter != m.rend(); ++iter)
        {
            float time = float(target - iter->first) / iter->second;
            stack.push_back(time);
            if (stack.size() >= 2)
            {
                if (stack.back() <= stack[stack.size() - 2])
                {
                    stack.pop_back();
                }
            }
        }
        return stack.size();
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}