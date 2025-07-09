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
    int optimizeBlockchain(vector<vector<int>> arr)
    {
        int maxS = arr[0][1];
        int N = arr[0][0];
        vector<vector<int>> calculatedArr;
        bool isRelate = false;
        for (int i = 1; i < arr.size(); i++)
        {
            vector<int> subArr = arr[i];
            if (subArr.size() == 1)
            {
                int d = subArr[0];
                isRelate = true;
                continue;
            }
            if (isRelate)
            {
                // first trans += second trans
                int idxFirstTrans = subArr[0] - 1;
                int idxSecondTrans = subArr[1] - 1;
                calculatedArr[idxFirstTrans][0] += calculatedArr[idxSecondTrans][0];
                calculatedArr[idxFirstTrans][1] += calculatedArr[idxSecondTrans][1];
            }
            else
            {
                calculatedArr.push_back(subArr);
            }
        }

        for (int i = 0; i < N - 1; i++)
        {
            for (int j = 0; j < N - 1 - i; j++)
            {
                if (calculatedArr[j][0] > calculatedArr[j + 1][0])
                {
                    swap(calculatedArr[j], calculatedArr[j + 1]);
                }
                else if (calculatedArr[j][0] == calculatedArr[j + 1][0])
                {
                    if (calculatedArr[j][1] > calculatedArr[j + 1][1])
                    {
                        swap(calculatedArr[j], calculatedArr[j + 1]);
                    }
                }
            }
        }

        vector<int> prefixSumFee;
        vector<int> prefixSumSize;

        int sumFee = 0;
        int sumSize = 0;
        for (auto el : calculatedArr)
        {
            sumFee += el[0];
            sumSize += el[1];
            prefixSumFee.push_back(sumFee);
            prefixSumSize.push_back(sumSize);
        }

        int left = 0, right = N - 1;
        int maxFee = 0;

        while (left <= right)
        {
            if (calculatedArr[right][1] > maxS)
            {
                right--;
                continue;
            }

            int calculatedSize = prefixSumSize[right] - ((left > 0) ? prefixSumSize[left - 1] : 0);
            int calculatedFee = prefixSumFee[right] - ((left > 0) ? prefixSumFee[left - 1] : 0);

            if (calculatedSize > maxS)
            {
                left++;
            }
            else
            {
                maxFee = (maxFee > calculatedFee) ? maxFee : calculatedFee;
                left = 0;
                right--;
            }
        }
        return maxFee;
    }
};

int main()
{
    Solution solution;
    vector<vector<int>> arr1{
        {3, 10},
        {10, 5},
        {7, 4},
        {3, 3},
        {1},
        {1, 2}};

    cout << "Result 1: " << solution.optimizeBlockchain(arr1) << endl;

    vector<vector<int>> arr2{
        {3, 10},
        {10, 5},
        {10, 5},
        {5, 3}};
    cout << "Result 2: " << solution.optimizeBlockchain(arr2) << endl;

    vector<vector<int>> arr3{
        {2, 10},
        {100, 11},
        {200, 12}};
    cout << "Result 3: " << solution.optimizeBlockchain(arr3) << endl;

    return 0;
    return 0;
}