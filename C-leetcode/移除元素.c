int removeElement(int* nums, int numsSize, int val){
int x1,x2=0;
for(int i=0;i<=numsSize-1;i++)
{
    if(nums[x1]!=val)
    {
        x2=x1;
        x2++;
    }
    x1++;
}
return x2;
}

