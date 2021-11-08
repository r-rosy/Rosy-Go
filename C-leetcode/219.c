typedef struct {
    int key;
    int val;
    UT_hash_handle hh;
}Hash;
bool containsNearbyDuplicate(int* nums, int numsSize, int k){
Hash* hash=NULL;
for(int i=0;i<=numsSize-1;i++)
{
    Hash* tmp;
    HASH_FIND_INT(hash,nums+i,tmp);
    if(tmp==NULL)
    {
        tmp=malloc(sizeof(Hash));
        tmp->key=nums[i];
        tmp->val=i;
        HASH_ADD_INT(hash,key,tmp);
    }else {
        if(tmp->val-i<=k&&tmp->val-i>=(-k))
        {
            return true;
        }else {
            Hash* tmp1=malloc(sizeof(Hash));
            tmp1->key=nums[i];
            tmp1->val=i;
            HASH_ADD(hh,hash,key,sizeof(int),tmp1);
            HASH_DEL(hash,tmp);
        }
    }
}
return false;
}