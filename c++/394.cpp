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
    string decodeString(string s)
    {

        // Example 1:
        // Input: s = "3[a]2[bc]"
        // Output: "aaabcbc"

        // Example 2:
        // Input: s = "3[a2[c]]"
        // Output: "accaccacc"

        // Example 3:
        // Input: s = "2[abc]3[cd]ef"
        // Output: "abcabccdcdcdef"

        // "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
        vector<string> stackS;
        vector<int> stackI;
        string kToRepeat = "";
        string strToRepeat = "";
        for (char c : s)
        {
            if (c >= '0' && c <= '9')
            {
                kToRepeat += c;
            }
            else if (c == '[')
            {
                stackS.push_back(strToRepeat);
                strToRepeat = "";
                stackI.push_back(stoi(kToRepeat));
                kToRepeat = "";
            }
            else if (c == ']')
            {
                string temp = "";
                for (int i = 0; i < stackI[stackI.size() - 1]; i++)
                {
                    temp += strToRepeat;
                }
                stackI.pop_back();
                strToRepeat = stackS[stackS.size() - 1] + temp;
                stackS.pop_back();
            }
            else
            {
                strToRepeat += c;
            }
        }
        return strToRepeat;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}