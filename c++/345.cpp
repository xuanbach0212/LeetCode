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
    string reverseVowels(string s)
    {
        unordered_set<char> vowels = {'a', 'i', 'u', 'e', 'o'};
        stack<char> stack;
        string result;

        for (char c : s)
        {
            auto it = vowels.find(tolower(c));
            if (it != vowels.end())
            {
                stack.push(c);
            }
        }

        for (char c : s)
        {
            auto it = vowels.find(tolower(c));
            if (it != vowels.end())
            {
                c = stack.top();
                stack.pop();
            }
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