int biresearch(int* nums,int numsSize,int target)
{
    int left=0,right=numsSize-1,mid=0;
    while(left<=right)
    {
        mid=(left+right)/2;
        if(nums[mid]==target)
        {
            return mid;
        }else if(nums[mid]<target)
        {
            left=mid+1;
        }else {
            right=mid-1;
        }
    }
    return -1;
}
int search(int* nums, int numsSize, int target){
return biresearch(nums,numsSize,target);
}

