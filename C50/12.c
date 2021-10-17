#include<stdio.h>
void calculation (int n,int* i) 
{
    if (n%2==0) 
    {
        n=n/2;
    } else {
        n=(3*n+1)/2;
    }
    *i=*i+1;
}
int main () 
{
    int a;
    int i=0;
    scanf("%d",&a);
    while(a!=1) 
    {
        calculation(a,&i);
    }
    printf("%d",i);
}