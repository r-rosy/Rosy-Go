void delete(int* nums,int n,int numsSize)
{
for(int i=n;i<=numsSize-2;i++)
{
    nums[i]=nums[i+1];
}
}


int removeDuplicates(int* nums, int numsSize){
    int n=0;
    for(int i=0;i<=numsSize-2;)
    {
        if(nums[i]!=nums[i+1])
        {
            i=i+1;
        }
        else{
            delete(nums,i,numsSize);
            numsSize--;
        }
    }
    return numsSize;
}

