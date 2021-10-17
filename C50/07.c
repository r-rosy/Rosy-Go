int product(int* array,int n)
{
    int pro=1;
    for(int i=0;i<=n-1;i++)
    {
        if(array[i]%2!=0)
        {
            pro=pro*array[i];
        }
    }
    return pro;
}