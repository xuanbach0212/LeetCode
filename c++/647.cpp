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
        int length = s.size();
        int count = length;
        int range = 2;
        string revertString = "";
        for (int i = length - 1; i >= 0; i--)
        {
            revertString += s[i];
        }

        while (range <= length)
        {

            for (int i = 0; i < length; i++)
            {
                int subStringLen = range;
                if (range + i > length)
                {
                    continue;
                }
                string subString = s.substr(i, subStringLen);
                // revertIndex = len - index - subStringLen
                string revertSubString = revertString.substr(length - i - subStringLen, subStringLen);
                if (subString == revertSubString)
                {
                    count += 1;
                }
            }
            range++;
        }
        return count;
    }
};

int main()
{
    Solution solution;
    string s = "abad";
    int result = solution.countSubstrings(s);

    cout << "Result: " << result << endl;
    return 0;
}