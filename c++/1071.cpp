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
    string gcdOfStrings(string str1, string str2)
    {
        string result;
        if (str1 + str2 != str2 + str1)
            return "";

        int a = str1.size();
        int b = str2.size();

        while (b != 0)
        {
            int r = a % b;
            a = b;
            b = r;
        }

        int i = 0;
        while (i < a)
        {
            result += str2[i];
            i++;
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