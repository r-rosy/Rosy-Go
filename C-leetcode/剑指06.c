/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* reversePrint(struct ListNode* head, int* returnSize){
    int size=0;
    for(struct ListNode* p=head;p!=NULL;p=p->next)
    {
        size++;
    }
    int arr[size];
    int i=0;
    for(struct ListNode* p=head;p!=NULL;p=p->next)
    {
        arr[size-i-1]=p->val;
        i++;
    }
    *(returnSize)=size;
    int* p=arr;
    return p;


}

