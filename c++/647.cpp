#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    int countSubstrings(string s)
    {
        int count = 0;
        for (int i = 0; i < s.size(); i++)
        {
            // odd case
            int l = i, r = i;
            while (l >= 0 && r < s.size() && s[l] == s[r])
            {
                count += 1;
                l -= 1;
                r += 1;
            }
            // even case
            l = i, r = i + 1;
            while (l >= 0 && r < s.size() && s[l] == s[r])
            {
                count += 1;
                l -= 1;
                r += 1;
            }
        }
        return count;
    }
};

int main()
{
    Solution solution;
    string s = "abc";
    int result = solution.countSubstrings(s);

    cout << "Result: " << result << endl;
    return 0;
}