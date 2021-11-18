/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
typedef struct {
    int key;
    int val;
    UT_hash_handle hh;
}Hash;

int* nextGreaterElement(int* nums1, int nums1Size, int* nums2, int nums2Size, int* returnSize){
int array[10000];
int top=-1;
Hash* hash=NULL;
for(int i=nums2Size-1;i>=0;i--)
{
    Hash* tmp=malloc(sizeof(Hash));
    tmp->key=nums2[i];
    if(top!=-1)
    {
        for(;;)
        {
            if(array[top]>=nums2[i])
            {
                break;
            }
            top--;
            if(top==-1)
            {
                break;
            }
        }
        if(top!=-1)
        {
        tmp->val=array[top];
        }else {
        tmp->val=-1;
        }
    }else {
        tmp->val=-1;
    }
    HASH_ADD_INT(hash,key,tmp);
    top++;
    array[top]=nums2[i];

}
int* arr=malloc(sizeof(int)*nums1Size);
for(int i=0;i<=nums1Size-1;i++)
{
    Hash* tmp;
    HASH_FIND_INT(hash,&nums1[i],tmp);
    arr[i]=tmp->val;
}
*returnSize=nums1Size;
return arr;
}