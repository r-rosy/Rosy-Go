/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* reverseList(struct ListNode* head){
    if(head==NULL)
    {
        return NULL;
    }
    if(head->next==NULL)
    {
        return head;
    }
    struct ListNode* p=head;
    for(;p->next!=NULL;)
    {
        struct ListNode* q=p->next;
        p->next->next=p;
        p->next=NULL;
        p=q;
    }
    return p;
}

