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
    int evalRPN(vector<string> &tokens)
    {
        // Input: tokens = ["2","1","+","3","*"]
        // Output: 9
        // Explanation: ((2 + 1) * 3) = 9

        // Input: tokens = ["4","13","5","/","+"]
        // Output: 6
        // Explanation: (4 + (13 / 5)) = 6
        stack<string> st;
        unordered_set<string> operators = {"+", "-", "*", "/"};
        for (string t : tokens)
        {
            string toPush;
            if (operators.count(t))
            {
                int val1 = stoi(st.top());
                st.pop();
                int val2 = stoi(st.top());
                st.pop();

                if (t == "+")
                    toPush = to_string(val2 + val1);
                else if (t == "-")
                    toPush = to_string(val2 - val1);
                else if (t == "*")
                    toPush = to_string(val2 * val1);
                else if (t == "/")
                    toPush = to_string(val2 / val1);
            }
            else
            {
                toPush = t;
            }
            st.push(toPush);
        }
        return stoi(st.top());
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}