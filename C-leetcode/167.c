/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* numbers, int numbersSize, int target, int* returnSize){
int* arr=(int*)malloc(sizeof(int)*2);
for(int i=0;i<=numbersSize-2;i++)
{
    for(int j=i+1;j<=numbersSize-1;j++)
    {
    if(numbers[i]+numbers[j]==target)
    {
        arr[0]=i+1;
        arr[1]=j+1;
        *returnSize=2;
        return arr;
    }
    }
}
*returnSize=2;
return arr;
}