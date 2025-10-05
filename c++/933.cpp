
#include <iostream>
#include <map>
#include <queue>
#include <set>
#include <stack>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

using namespace std;

class RecentConter {
public:
  queue<int> q;
  RecentConter() {}
  int ping(int t) {
    q.push(t);
    while (q.front() < t - 3000) {
      q.pop();
    }
    return q.size();
  }
};

int main() {
  RecentConter rc;
  vector<int> inputs = {1, 100, 3001, 3002};
  for (int t : inputs) {
    cout << rc.ping(t) << " ";
  }
  cout << endl;
}
