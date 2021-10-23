#include<stdio.h>
int main()
{
    int a,b;
    scanf("%d %d",&a,&b);
    if((a<0||a>100)||(b<0||b>100))
    {
        printf("it is a error");
        return 0;
    }
    if(a>60&&b>60)
    {
        printf("it is pass");
        return 0;
    }else
    {
        printf("it is not pass");
        return 0;
    }
}