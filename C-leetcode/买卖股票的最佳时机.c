int com(const void* a,const void* b)
{
    return *(int*)a-*(int*)b;
}
int maxProfit(int* prices, int pricesSize){
if(pricesSize==1)
{
    return 0;
}
int n=pricesSize;
int num=(n*(n-1))/2; 
int array[num];
int t=0;
for(int i=0;i<=pricesSize-2;i++)
{
    for(int j=i+1;j<=pricesSize-1;j++)
    {
        array[t]=prices[j]-prices[i];
        t++;
    }
}
qsort(array,num,sizeof(int),com);
if(array[num-1]>0)
{
    return array[num-1];
}else {
    return 0;
}
}
