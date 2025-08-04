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
    string reverseWords(string s)
    {
        // s = "   the sky is blue"
        // "blue is sky the"
        // s = "  hello world "
        // "world hello"
        // s = "t "
        // "t"
        string result;
        int left = 0;
        int right = 0;
        vector<string> v;

        while (left < s.size())
        {
            if (s[left] == ' ')
            {
                left++;
                continue;
            }

            right = left;
            while (right < s.size() && s[right] != ' ')
            {
                right++;
            }

            string word(s.begin() + left, s.begin() + right);
            v.push_back(word);
            left = right;
        }

        for (int i = v.size() - 1; i >= 0; i--)
        {
            result += v[i];
            if (i != 0)
            {
                result += " ";
            }
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