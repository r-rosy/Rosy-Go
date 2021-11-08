typedef struct {
    char key;
    char val;
    UT_hash_handle hh;
}Hash;
bool is(char* s,char* t){
if(strlen(s)!=strlen(t))
{
    return false;
}
Hash* hash=NULL;
for(int i=0;i<=strlen(s)-1;i++)
{
    Hash* tmp;
    HASH_FIND(hh,hash,s+i,sizeof(char),tmp);
    if(tmp==NULL)
    {
        tmp=malloc(sizeof(Hash));
        tmp->key=s[i];
        tmp->val=t[i];
        HASH_ADD(hh,hash,key,sizeof(char),tmp);
    }else {
        if(tmp->val!=t[i])
        {
            return false;
        }
    }
}
return true;
}
bool isIsomorphic(char * s, char * t){
if(is(s,t)==true)
{
    if(is(t,s)==true)
    {
        return true;
    }
}
return false;
}