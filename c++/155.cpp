#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>
#include <unordered_set>
#include <stack>
#include <unordered_map>

using namespace std;

class MinStack
{
private:
    stack<int> s;
    stack<int> minS;

public:
    MinStack()
    {
    }
    // [-3, 0, -2] s
    // [-3, -3, -3] minS
    void push(int val)
    {
        s.push(val);
        if (!minS.empty())
        {
            val = (minS.top() >= val) ? val : minS.top();
        }
        minS.push(val);
    }

    void pop()
    {
        s.pop();
        minS.pop();
    }

    int top()
    {
        return s.top();
    }

    int getMin()
    {
        return minS.top();
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}