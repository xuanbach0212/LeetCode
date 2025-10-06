#include <iostream>
#include <list>
#include <map>
#include <queue>
#include <set>
#include <stack>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
  // Input: head = [1,2,3,4,5]
  // Output: [1,3,5,2,4]
  ListNode *oddEvenList(ListNode *head) {
    if (!head || !head->next || !head->next->next) {
      return head;
    }
    ListNode *odd = head;
    ListNode *even = head->next;
    ListNode *evenHead = even;
    while (even && even->next) {
      odd->next = even->next;
      odd = odd->next;
      even->next = odd->next;
      even = even->next;
    }
    odd->next = evenHead;
    return head;
  }
};

ListNode *createList(const vector<int> &vals) {
  if (vals.empty())
    return nullptr;
  ListNode *head = new ListNode(vals[0]);
  ListNode *curr = head;
  for (size_t i = 1; i < vals.size(); ++i) {
    curr->next = new ListNode(vals[i]);
    curr = curr->next;
  }
  return head;
}

void printList(ListNode *head) {
  while (head) {
    cout << head->val;
    if (head->next)
      cout << " -> ";
    head = head->next;
  }
  cout << '\n';
}

void freeList(ListNode *head) {
  while (head) {
    ListNode *next = head->next;
    delete head;
    head = next;
  }
}

int main() {
  Solution solution;
  vector<vector<int>> tests = {
      {1, 2, 3, 4, 5},
      {2, 1, 3, 5, 6, 4, 7},
  };

  for (auto &arr : tests) {
    ListNode *head = createList(arr);
    cout << "Before: ";
    printList(head);

    head = solution.oddEvenList(head);

    cout << "After:  ";
    printList(head);
    cout << "-----------------------\n";

    freeList(head);
  }

  return 0;
}
