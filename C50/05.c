#include<stdio.h>
int main()
{
    printf("%d ",2);
    for(int i=3;i<=100;i++)
    {
        int t=0;
        for(int j=2;j<=i-1;j++)
        {
            if(i%j==0)
            {
                break;
            }
            if(j==i-1)
            {
                t=1;
            }
        }
        if(t!=0)
        {
            printf("%d ",i);
        }
    }
}