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
        int currentSum = 0;
        int different;
        // make a map like prefixSumMap(currentSum, count to this sum)
        unordered_map<int, int>
            prefixSumMap;
        prefixSumMap.insert({0, 1}); // for base case
        for (int num : nums)
        {
            currentSum += num;
            different = currentSum - k;
            count += prefixSumMap[different]; // add the sub array count
            prefixSumMap[currentSum]++;
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