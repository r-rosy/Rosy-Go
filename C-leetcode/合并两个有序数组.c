void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n){
int j=0;
for(int i=m;i<m+n;i++)
{
    nums1[i]=nums2[j];
    j++;
}
//选择排序
for(int i=0;i<nums1Size-1;i++)
{
    for(int t=i+1;t<nums1Size;t++)
    {
        if(nums1[i]>nums1[t])
        {
            int n=nums1[i];
            nums1[i]=nums1[t];
            nums1[t]=n;
        }
    }
}
}


