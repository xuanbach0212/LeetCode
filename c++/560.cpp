#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    int subarraySum(vector<int> &nums, int k)
    {
        int count = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            int right = i;
            int sum = 0;
            while (right < nums.size())
            {
                sum += nums[right];
                if (sum == k)
                {
                    count += 1;
                }
                right += 1;
            }
        }
        return count;
    }
};

int main()
{
    Solution solution;
    string result = "abc";

    cout << "Result: " << result << endl;
    return 0;
}