int strStr(char * haystack, char * needle){
if(strlen(needle)==0)
{
    return 0;
}
if(strlen(haystack)==0)
{
    return -1;
}
for(int i=0;i<=strlen(haystack)-1;i++)
{
    if(haystack[i]==needle[0])
    {
        int t=0;
        for(int j=0;j<=strlen(needle)-1;j++)
        {
            if(i+t>=strlen(haystack))
            {
                return -1;
            }
            if(haystack[i+t]!=needle[j])
            {
                break;
            }
            t++;
        }
        if(t==strlen(needle))
        {
            return i;
        }
    }
}
return -1;
}