/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int com(const void* a,const void* b)
{
    return *(int*)a-*(int*)b;
}
int* getLeastNumbers(int* arr, int arrSize, int k, int* returnSize){
qsort(arr,arrSize,sizeof(int),com);
int* ar=(int*)malloc(sizeof(int)*k);
for(int i=0;i<=k-1;i++)
{
    ar[i]=arr[i];
}
*(returnSize)=k;
return ar;
}

