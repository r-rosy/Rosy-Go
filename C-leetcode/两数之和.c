int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    int* array=(int*)malloc(sizeof(int)*2);
    for(int i=0;i<=numsSize-2;i++)
    {
        for(int j=i+1;j<=numsSize-1;j++)
        {
            if(nums[i]+nums[j]==target)
            {
                array[0]=i;
                array[1]=j;
                *returnSize =2;
                return array;
            }
        }
    }
    return array;
}
