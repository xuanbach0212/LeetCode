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
    int largestAltitude(vector<int> &gain)
    {
        // Input: gain = [-5,1,5,0,-7]
        // Output: 1
        // Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
        int maxAltitude = 0;
        int previousAltitude = 0;
        for (int i = 0; i < gain.size(); i++)
        {
            previousAltitude += gain[i];
            maxAltitude = (maxAltitude > previousAltitude) ? maxAltitude : previousAltitude;
        }
        return maxAltitude;
    }
};

int main()
{
    // Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}