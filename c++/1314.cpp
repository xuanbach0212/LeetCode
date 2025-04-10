#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    vector<vector<int>> matrixBlockSum(vector<vector<int>> &mat, int k)
    {
        vector<vector<int>> answer(mat.size(), vector<int>(mat[0].size(), 0));
        for (int i = 0; i < mat.size(); i++)
        {
            int leftR = i - k;
            int rightR = i + k;

            if (leftR < 0)
                leftR = 0;
            if (rightR > mat.size() - 1)
                rightR = mat.size() - 1;

            for (int j = 0; j < mat[i].size(); j++)
            {
                int leftC = j - k;
                int rightC = j + k;

                if (leftC < 0)
                    leftC = 0;
                if (rightC > mat[i].size() - 1)
                    rightC = mat[i].size() - 1;

                int sum = 0;
                for (int r = leftR; r <= rightR; r++)
                {
                    for (int c = leftC; c <= rightC; c++)
                    {
                        sum += mat[r][c];
                    }
                }

                answer[i][j] = sum;
            }
        }
        return answer;
    }
};

int main()
{
    Solution solution;
    vector<int> nums = {1, 1, 2, 2, 3, 4, 4, 5};

    return 0;
}