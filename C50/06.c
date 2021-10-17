int jd(int x)
{
    if(x<0)
    {
        return -x;
    }else{
        return x;
    }
    return 0;
}
void selecetnum(int* arr,int n)
{
    for(int i=0;i<=n-2;i++)
    {
      for(int j=i+1;j<=n-1;j++)
      {
          if(jd(arr[i])>jd(arr[j]))
          {
              int n=arr[j];
              arr[j]=arr[i];
              arr[i]=n;
          }
      }
    }
}