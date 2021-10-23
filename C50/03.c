#include<stdio.h>
int main()
{
    for(int i=0;i<=2;i++)
    {
        for(int j=0;j<i;j++)
        {
            printf(" ");
        }
        for(int m=0;m<5-2*i;m++)
        {
            printf("*");
        }
        printf("\n");
    }
    printf(" ***\n");
    printf("*****");
}