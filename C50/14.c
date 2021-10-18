int* change(int* array,int n)
{
int arr[n];
for(int i=0;i<=n-1;i++)
{
    arr[i]=0;
}
for(int i=0;i<=n-1;i++)
{
arr[i]=array[i];
}
for(int i=0;i<=n-2;i++)
{
    for(int j=i+1;j<=n-1;j++)
    {
        if(arr[i]>arr[j])
        {
            int p=arr[i];
            arr[i]=arr[j];
            arr[j]=p;
        }
    }
}
int i=0;
for(;i<=n-1;i++)
{
    if(array[i]==arr[0])
    {
        break;
    }
}
int t=array[0];
array[0]=array[i];
array[i]=t;
return array;
}