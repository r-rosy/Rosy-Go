int com(const void* a,const void* b)
{
    return (*(int*)a-*(int*)b);
}
int singleNumber(int* nums, int numsSize){
qsort(nums,numsSize,sizeof(int),com);
for(int i=0;i<=numsSize-2;)
{
    int n=1;
    if(nums[0]!=nums[1])
    {
        return nums[0];
    }
    for(int j=i+1;j<=numsSize-1;j++)
    {
        if(nums[i]!=nums[j])
        {
            if(n==1)
            {
                return nums[i];
            }
            i=j;
            break;
        }
        n++;
    }
}
return nums[numsSize-1];
}

