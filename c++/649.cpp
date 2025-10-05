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

class Solution {
public:
  string predictPartyVictory(string senate) {
    // [ D D R R R ]
    queue<int> rQueue;
    queue<int> dQueue;
    for (int i = 0; i < senate.size(); i++) {
      if (senate[i] == 'R') {
        rQueue.push(i);
      } else {
        dQueue.push(i);
      }
    }
    while (rQueue.size() != 0 && dQueue.size() != 0) {
      int rTurn = rQueue.front();
      int dTurn = dQueue.front();
      rQueue.pop();
      dQueue.pop();

      if (rTurn < dTurn) {
        rQueue.push(rTurn + senate.size());
      } else {
        dQueue.push(dTurn + senate.size());
      }
    }
    return rQueue.size() > dQueue.size() ? "Radiant" : "Dire";
  }
};

int main() {
  Solution solution;
  vector<string> s{"RD", "RDD", "DRRD", "DDRRR"};
  for (string el : s) {
    string result = solution.predictPartyVictory(el);
    cout << "with " << el << " result: " << result << endl;
  }
}
