/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2){
if(l1==NULL)
{
    return l2;
}else if(l2==NULL)
{
    return l1;
}
struct ListNode* p=l1;
for(;p->next!=NULL;p=p->next)
{
    
}
p->next=l2;
for(struct ListNode* m=l1;m->next!=NULL;m=m->next)
{
    for(struct ListNode* n=m->next;n!=NULL;n=n->next)
    {
        if(m->val>n->val)
        {
            int t=m->val;
            m->val=n->val;
            n->val=t;
        }
    }
}
return l1;
}

