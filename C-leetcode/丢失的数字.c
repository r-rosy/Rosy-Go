int com(const void* a,const void* b)
{
    return (*(int*)a-*(int*)b);
}
int missingNumber(int* nums, int numsSize){
qsort(nums,numsSize,sizeof(int),com);
for(int i=0;i<=numsSize-1;i++)
{
    if(nums[i]!=i)
    {
        return i;
    }
}
return numsSize;
}

