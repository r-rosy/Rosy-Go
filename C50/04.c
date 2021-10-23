#include<stdio.h>
int main()
{
    int a,b;
    scanf("%d %d",&a,&b);
    int n=a>b?b:a;
    int m=a>b?a:b;
    int t;
    for(int i=1;i<=n;i++)
    {
        if(n%i==0)
        {
            if(m%i==0)
            {
                t=i;
            }
        }
    }
    printf("%d",t);
}