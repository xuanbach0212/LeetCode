#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>
#include <unordered_set>

using namespace std;

class Solution
{
public:
    bool isValidSudoku(vector<vector<char>> &board)
    {
        // make hash map with hash set for value
        // row and column key is char
        unordered_map<int, unordered_set<char>> row;
        unordered_map<int, unordered_set<char>> column;
        unordered_map<int, unordered_set<char>> square; // square key is tuple (r/3, c/3)

        for (int r = 0; r < 9; r++)
        {
            for (int c = 0; c < 9; c++)
            {
                // if empty continue
                if (board[r][c] == '.')
                    continue;
                // check if row, column, square not duplicate
                int squareKey = (r / 3) * 3 + (c / 3);
                if (row[r].find(board[r][c]) != row[r].end() || column[c].find(board[r][c]) != column[c].end() || square[squareKey].find(board[r][c]) != square[squareKey].end())
                {
                    return false;
                }
                row[r].insert(board[r][c]);
                column[c].insert(board[r][c]);
                square[squareKey].insert(board[r][c]);
            }
        }
        return true;
    }
};

int main()
{
    Solution solution;
    vector<vector<char>> board = {
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '8', '.', '.', '.', '3', '.'},
        {'4', '.', '7', '.', '.', '.', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'8', '5', '9', '7', '.', '.', '.', '.', '.'},
        {'6', '7', '9', '.', '.', '.', '.', '.', '.'},
        {'1', '9', '3', '4', '2', '8', '5', '6', '7'}};

    bool result = solution.isValidSudoku(board);
    cout << result << endl;
}