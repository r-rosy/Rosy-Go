int removeElement(int* nums, int numsSize, int val){
int x2=0;
for(int x1=0;x1<=numsSize-1;x1++)
{
    if(nums[x1]!=val)
    {
        nums[x2]=nums[x1];
        x2++;
    }
}
numsSize=x2;
return x2;
}

