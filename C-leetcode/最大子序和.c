int max(int a,int b)
{
    return a>b?a:b;
}
int SumSelect(int* nums,int n,int numsSize){
if(n==0)
{
    return nums[0];
}else{
    return max(nums[n],SumSelect(nums,n-1,numsSize)+nums[n]);
}
}
int com(const void* a,const void* b)
{
    return *(int*)a-*(int*)b;
}
int maxSubArray(int* nums, int numsSize){
int n=numsSize;
int* result=(int*)malloc(sizeof(int)*n);
for(int i=0;i<=n-1;i++)
{
result[i]=SumSelect(nums,i,numsSize);
}
qsort(result,n,sizeof(int),com);
return result[n-1];
}