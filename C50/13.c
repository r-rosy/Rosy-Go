#include<stdio.h>
int main()
{
    int T;
    scanf("%d",&T);
    for(int i=1;i<=T;i++)
    {
        int a,b,c;
        scanf("%d %d %d",&a,&b,&c);
        if(a+b>c)
        {
            printf("Case #%d: true\n",i);
        }else {
            printf("Case #%d: false\n",i);
        }
        
    }
}