int com(const void* a,const void* b)
{
    return (int)*(char*)a-(int)*(char*)b;
}
bool isAnagram(char * s, char * t){
if(strlen(s)!=strlen(t))
{
    return false;
}
qsort(s,strlen(s),sizeof(char),com);
qsort(t,strlen(t),sizeof(char),com);
for(int i=0;i<=strlen(t)-1;i++)
{
    if(s[i]!=t[i])
    {
        return false;
    }
}
return true;
}