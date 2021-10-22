/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
for(struct ListNode* p=headA;p!=NULL;p=p->next)
{
    for(struct ListNode* q=headB;q!=NULL;q=q->next)
    {
        if(p==q)
        {
            return p;
        }
    }
}
return NULL;
}

