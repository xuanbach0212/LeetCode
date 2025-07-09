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
    int largestRectangleArea(vector<int> &heights)
    {
        // [2,1,5,6,2,3] = 10
        // [4,2,0,3,1,4,3,4]
        // [0,9]
        // [2,4]
        // [2,1,2]
        int maxArea = 0;
        stack<pair<int, int>> s;
        for (int i = 0; i < heights.size(); ++i)
        {
            int start = i;
            while (!s.empty() && s.top().second > heights[i])
            {
                int index = s.top().first;
                int height = s.top().second;
                s.pop();
                int area = height * (i - index);
                maxArea = max(maxArea, area);
                start = index;
            }

            s.push({start, heights[i]});
        }

        int n = heights.size();
        while (!s.empty())
        {
            int index = s.top().first;
            int height = s.top().second;
            s.pop();
            int area = height * (n - index);
            maxArea = max(maxArea, area);
        }

        return maxArea;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}