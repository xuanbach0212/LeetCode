#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>
#include <stack>

using namespace std;

class Solution
{
public:
    bool isValid(string s)
    {
        unordered_map<char, char> m{
            {'}', '{'},
            {']', '['},
            {')', '('},
        };
        stack<char> stackChar;
        for (char c : s)
        {
            if (m.count(c))
            {
                if (!stackChar.empty() && stackChar.top() == m[c])
                {
                    stackChar.pop();
                }
                else
                {
                    return false;
                }
            }
            else
            {
                stackChar.push(c);
            }
        }
        return stackChar.empty();
    }
};

int main()
{
    Solution solution;
    string result = "abc";

    cout << "Result: " << result << endl;
    return 0;
}