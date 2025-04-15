#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class NumMatrix
{
private:
    vector<vector<int>> prefixSumMatrix;

public:
    NumMatrix(vector<vector<int>> &matrix)
    {
        prefixSumMatrix.resize(matrix.size(), vector<int>(matrix[0].size(), 0));
        for (int i = 0; i < matrix.size(); i++)
        {
            int currentSumRow = 0;
            for (int j = 0; j < matrix[i].size(); j++)
            {
                currentSumRow += matrix[i][j];
                prefixSumMatrix[i][j] = currentSumRow;
            }
        }
    }

    int sumRegion(int row1, int col1, int row2, int col2)
    {
        int currRow = row1;
        int sums = 0;
        for (int row = row1; row <= row2; row++)
        {
            vector<int> prefixRow = prefixSumMatrix[currRow];
            int left = (col1 > 0) ? prefixRow[col1 - 1] : 0;
            int right = prefixRow[col2];
            sums += (right - left);
            currRow++;
        }
        return sums;
    }
};

/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix* obj = new NumMatrix(matrix);
 * int param_1 = obj->sumRegion(row1,col1,row2,col2);
 */