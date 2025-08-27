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
    string removeStars(string s)
    {
        vector<char> stack;
        for (char c : s)
        {
            if (c != '*')
            {
                stack.push_back(c);
            }
            else
            {
                stack.pop_back();
            }
        }
        string result = "";
        for (char c : stack)
        {
            result += c;
        }
        return result;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}