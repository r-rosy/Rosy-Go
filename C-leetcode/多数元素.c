int com(const void* a,const void* b)
{
    return (*(int*)a-*(int*)b);
}
int majorityElement(int* nums, int numsSize){
qsort(nums,numsSize,sizeof(int),com);
for(int i=0;i<numsSize-1;)
{
    int n=1;
    for(int j=i+1;j<=numsSize-1;j++)
    {
        if(nums[i]==nums[j])
        {
            n++;
        }else{
        if(n>(numsSize/2))
            {
                return nums[i];
            }
            i+=n;
            break;
        }
    }
}
return 0;
}

