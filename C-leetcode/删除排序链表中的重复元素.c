/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* deleteDuplicates(struct ListNode* head){
if(head==NULL||head->next==NULL)
{
    return head;
}
struct ListNode* q=NULL;
for(struct ListNode* p=head;p!=NULL&&p->next!=NULL;)
{
    q=p->next;
    if(p->val==q->val)
    {
        p->next=q->next;
        q->next=NULL;
    }else{
    p=p->next;
    }
}
return head;
}

