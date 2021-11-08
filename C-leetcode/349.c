/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
typedef struct {
    int key;
    UT_hash_handle hh;
}Hash;
int* intersection(int* nums1, int nums1Size, int* nums2, int nums2Size, int* returnSize){
Hash* hash=NULL;
Hash* hashh=NULL;
for(int i=0;i<=nums1Size-1;i++)
{
    Hash* tmp;
    HASH_FIND_INT(hash,nums1+i,tmp);
    if(tmp==NULL)
    {
        tmp=malloc(sizeof(Hash));
        tmp->key=nums1[i];
        HASH_ADD_INT(hash,key,tmp);
    }
}
int array[nums1Size];
int n=0;
for(int i=0;i<=nums2Size-1;i++)
{
    Hash* tmp;
    HASH_FIND_INT(hash,nums2+i,tmp);
    Hash* tmpp;
    HASH_FIND_INT(hashh,nums2+i,tmpp);
    if(tmpp==NULL)
    {
        tmpp=malloc(sizeof(Hash));
        tmpp->key=nums2[i];
        HASH_ADD_INT(hashh,key,tmpp);
        if(tmp!=NULL){
        array[n]=nums2[i];
        n++;
        }
    }
}
int* arr=malloc(sizeof(int)*n);
for(int i=0;i<=n-1;i++)
{
    arr[i]=array[i];
}
*returnSize=n;
return arr;
}