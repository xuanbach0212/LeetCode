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
  ListNode *deleteMiddle(ListNode *head) {
    if (!head || !head->next) {
      delete head;
      return nullptr;
    }

    int count = 0;
    ListNode *temp = head;
    while (temp) {
      ++count;
      temp = temp->next;
    }

    int neededIdx = count / 2;

    ListNode *prev = head;
    for (int i = 0; i < neededIdx - 1; i++) {
      prev = prev->next;
    }

    ListNode *toDel = prev->next;
    prev->next = toDel->next;
    toDel->next = nullptr;
    delete toDel;
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
      {1, 3, 4, 7, 1, 2, 6}, {1, 2, 3, 4}, {2, 1}, {5}};

  for (auto &arr : tests) {
    ListNode *head = createList(arr);
    cout << "Before: ";
    printList(head);

    head = solution.deleteMiddle(head);

    cout << "After:  ";
    printList(head);
    cout << "-----------------------\n";

    freeList(head);
  }

  return 0;
}
