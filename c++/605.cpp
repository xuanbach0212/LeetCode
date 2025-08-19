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
    bool canPlaceFlowers(vector<int> &flowerbed, int n)
    {
        for (int i = 0; i < flowerbed.size(); i++)
        {
            int left = (i > 0) ? flowerbed[i - 1] : 0;
            int right = (i < flowerbed.size() - 1) ? flowerbed[i + 1] : 0;
            int middle = flowerbed[i];
            if (middle == 0 && left == 0 && right == 0)
            {
                n--;
                flowerbed[i] = 1;
            }
            if (n == 0)
                return true;
        }
        return n <= 0;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}