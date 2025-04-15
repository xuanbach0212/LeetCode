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
        // add extra column and row 0,0 for not out of bound case
        // [0, 0,  0,  0]
        // [0, 3,  3,  4]
        // [0, 8, 14, 18]
        // [0, 9, 17, 21]
        prefixSumMatrix = vector<vector<int>>(matrix.size() + 1, vector<int>(matrix[0].size() + 1, 0));
        for (int i = 0; i < matrix.size(); i++)
        {
            int prefixSumRow = 0;
            for (int j = 0; j < matrix[i].size(); j++)
            {
                // prefix rectangle = prefix row + prefix above rectangle
                prefixSumRow += matrix[i][j];
                int prefixAboveRectangle = prefixSumMatrix[i][j + 1];
                prefixSumMatrix[i + 1][j + 1] = prefixSumRow + prefixAboveRectangle;
            }
        }
    }

    int sumRegion(int row1, int col1, int row2, int col2)
    {
        row1++, col1++, row2++, col2++;
        // sum = bottom right - left column - above rectangle + top left
        int bottomRight = prefixSumMatrix[row2][col2];
        int left = prefixSumMatrix[row2][col1 - 1];
        int above = prefixSumMatrix[row1 - 1][col2];
        int topLeft = prefixSumMatrix[row1 - 1][col1 - 1];

        return bottomRight - left - above + topLeft;
    }
};

/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix* obj = new NumMatrix(matrix);
 * int param_1 = obj->sumRegion(row1,col1,row2,col2);
 */