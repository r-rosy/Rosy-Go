#include<stdio.h>
#include<malloc.h>
#include<stdlib.h>
int com(const void*a,const void* b)
{
    return *(int*)a-*(int*)b;
}
int main()
{
    int a,b,c;
    int arry[3];
    int* arr=arry;
    scanf("%d %d %d",&a,&b,&c);
    arr[0]=a;
    arr[1]=b;
    arr[2]=c;
    qsort(arr,3,sizeof(int),com);
    if(arr[0]+arr[1]<=arr[2])
    {
        printf("%d",0);
        return 0;
    }
    if(arr[0]==arr[1]&&arr[1]==arr[2])
    {
        printf("%d",1);
        return 0;
    }
    if((arr[0]==arr[1]&&arr[1]!=arr[2])||(arr[1]==arr[2]&&arr[0]!=arr[1]))
    {
        printf("%d",2);
        return 0;
    }
    printf("%d",3);
}
