#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    int maxArea(vector<int> &height)
    {
        // max (distance of 2 height x height of smaller one)
        int left = 0, right = height.size() - 1;
        int biggest = 0;
        while (left < right)
        {
            int area = (right - left) * ((height[left] > height[right]) ? height[right] : height[left]);
            biggest = (biggest > area) ? biggest : area;

            if (height[left] > height[right])
            {
                right--;
            }
            else
            {
                left++;
            }
        }
        return biggest;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}