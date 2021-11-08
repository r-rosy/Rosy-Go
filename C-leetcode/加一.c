/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* addone(int* digits,int i)
{
if(digits[i]!=9)
{
digits[i]++;
return digits;
}else{
    digits[i]=0;
    addone(digits,i-1);
}
return digits;
}
int* plusOne(int* digits, int digitsSize, int* returnSize){
*returnSize=digitsSize;
if(digits[0]==9)
{
    (*returnSize)=(*returnSize)+1;
    digits[digitsSize]=0;
    for(int i=0;i<=digitsSize-1;i++)
    {
        digits[i+1]=digits[i];
        digits[0]=0;
    }
}
if(digits[digitsSize-1]!=9)
{
    digits[digitsSize-1]++;
    return digits;
}
return addone(digits,*returnSize-1);
}