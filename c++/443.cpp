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
    int compress(vector<char> &chars)
    {
        // ["a","a","b","b","c","c","c"]
        string s = "";
        int left = 0;
        int right = 0;
        char tmp = chars[0];
        while (left < chars.size())
        {
            if (right == chars.size() || tmp != chars[right])
            {
                s += chars[left];
                if (right - left != 1)
                {
                    for (char c : to_string(right - left))
                    {
                        s += c;
                    }
                }
                if (right == chars.size())
                {
                    break;
                }
                tmp = chars[right];
                left = right;
            }
            else
            {
                right++;
                continue;
            }
        }
        chars.clear();
        for (char c : s)
        {
            chars.push_back(c);
        }
        return s.size();
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}