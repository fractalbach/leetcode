// https://leetcode.com/problems/merge-two-sorted-lists/

struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
  ListNode *mergeTwoLists(ListNode *list1, ListNode *list2) {

    ListNode dummy = ListNode{};
    ListNode *head = &dummy;
    ListNode *curr = head;

    while (list1 != nullptr || list2 != nullptr) {

      if (list1 == nullptr) {
        curr->next = list2;
        break;
      }
      if (list2 == nullptr) {
        curr->next = list1;
        break;
      }

      if (list1->val < list2->val) {
        curr->next = list1;
        curr = list1;
        list1 = list1->next;
      } else {
        curr->next = list2;
        curr = list2;
        list2 = list2->next;
      }

    }
    return head->next;
  }
};
