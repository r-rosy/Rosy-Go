char check(char c)
{
    switch(c)
    {
        case ')':return '(';
        case ']':return '[';
        case '}':return '{';
        default :return ' ';
    }
}
bool isValid(char * s){
char array[1000000];
int top=-1;
if(strlen(s)==1)
{
    return false;
}
for(int i=0;i<=strlen(s)-1;i++)
{
    if(s[i]=='('||s[i]=='['||s[i]=='{')
    {
        top++;
        array[top]=s[i];
    }else {
        if(i==0||top==-1)
        {
            return false;
        }
        if(array[top]==check(s[i]))
        {
            top--;
        }else{
            return false;
        }
    }
}
if(top==-1)
{
    return true;
}else {
    return false;
}
return true;
}