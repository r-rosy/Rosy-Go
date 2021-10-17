bool Isreserve(int* num,long len)
{
    for(int i=0;i<=(len/2);i++)
    {
        if(num[i]!=num[len-i-1])
        {
            return false;
        }
    }
    return true;
}
bool isPalindrome(int x){
    if(x<0)
    {
        return false;
    }else if(x==0){
        return true;
    }else{
        int i=x;
        long t=0;
        while(i)
        {
         i=x/10;
         t++;
        }
        int* array=(int*)malloc(sizeof(int)*t);
        for(int i=0;i<=t-1;t++)
        {
            array[i]=(x/(10^i))%10;
        }
        return Isreserve(array,t);
    }

}
