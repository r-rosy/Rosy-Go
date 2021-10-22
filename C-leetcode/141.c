/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool hasCycle(struct ListNode *head) {
    struct ListNode* a=head;
    if(head==NULL||head->next==NULL)
    {
        return false;
    }
    struct ListNode* b=head->next;
    for(;b->next!=NULL&&b->next->next!=NULL;a=a->next)
    {
        if(a==b)
        {
            return true;
        }
        b=b->next->next;
    }
    return false;
}

