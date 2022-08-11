/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {

        int sum = 0;
        int carry = 0;
        int l1val = 0;
        int l2val = 0;
        ListNode* l3 = new ListNode();
        ListNode* head = l3;
        
        while (true) {
        
            if (l1 != nullptr) {
                l1val = l1->val;
                l1 = l1->next;
            } else {
                l1val = 0;
            }
            
            if (l2 != nullptr) {
                l2val = l2->val;
                l2 = l2->next;
            } else {
                l2val = 0;
            }
            
            sum = l1val + l2val + carry;
            carry = sum >= 10;
            sum = sum % 10;
            
            l3->val = sum;
            
            if ((l1 != nullptr) || (l2 != nullptr) || (carry == 1)) {
            	l3->next = new ListNode();
        		l3 = l3->next;
            } else {
            	break;
            }
        }
        
        return head;
    }
};
