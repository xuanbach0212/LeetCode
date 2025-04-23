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
    long long calculateScore(vector<string> &instructions, vector<int> &values)
    {
        // ["jump","add","add","jump","jump","jump"]
        // [3,5,12,-1,-3,-5]

        vector<bool> v(instructions.size(), false);
        long long score = 0;
        int addIndexNum = 1;
        int i = 0;
        while (i < instructions.size() && i >= 0)
        {
            if (v[i] == false)
            {
                v[i] = true;
            }
            else
            {
                break;
            }
            if (instructions[i] == "jump")
            {
                addIndexNum = values[i];
            }
            else
            {
                score += values[i];
                addIndexNum = 1;
            }

            i += addIndexNum;
        }
        return score;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}