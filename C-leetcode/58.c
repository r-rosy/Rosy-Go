int lengthOfLastWord(char * s){
int num=0;
for(int i=strlen(s)-1;i>=0;i--)
{
    if(s[i]!=' ')
    {
        for(int j=i;j>=0;j--)
        {
            if(s[j]==' ')
            {
                break;
            }
            num++;
        }
        break;
    }
}
return num;
}
