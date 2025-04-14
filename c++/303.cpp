#include <iostream>
#include <set>
#include <string>
#include <vector>
#include <map>

using namespace std;

class NumArray
{
private:
    vector<int> nums;

public:
    NumArray(vector<int> &nums)
    {
        this->nums = nums;
    }

    int sumRange(int left, int right)
    {
        int sums = 0;
        for (int i = left; i <= right; i++)
        {
            sums += this->nums[i];
        }
        return sums;
    }
};

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray* obj = new NumArray(nums);
 * int param_1 = obj->sumRange(left,right);
 */

int main()
{
    vector<int> data = {-2, 0, 3, -5, 2, -1};

    NumArray *obj = new NumArray(data);

    cout << "sumRange(0, 2): " << obj->sumRange(0, 2) << endl; // Kết quả mong đợi: 1
    cout << "sumRange(2, 5): " << obj->sumRange(2, 5) << endl; // Kết quả mong đợi: -1
    cout << "sumRange(0, 5): " << obj->sumRange(0, 5) << endl; // Kết quả mong đợi: -3

    delete obj;
    return 0;
}