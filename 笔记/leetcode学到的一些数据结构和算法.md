# C语言内置函数

***

```c
void qsort(void *base, size_t nitems, size_t size, int (*compar)(const void *, const void*))
```

<u>实现对数组的自动排序</u>

### 选择排序的实现（C++）

```cpp
#include<iostream>
using namespace std;
 
const int BUFFSIZE = 6;
void Selection_Sort(int Arr[]);

int main()
{
	int Arr[BUFFSIZE] = { 5,4,89,10,7,100}; //待排序元素
	Selection_Sort(Arr);

	return 0;
}

// 选择排序，从大到小排序
void Selection_Sort(int Arr[])
{
	for (int i = 0; i < BUFFSIZE-1 ; i++)  
	{
		int index = i;
		for (int j = i + 1; j < BUFFSIZE; j++)
		{			
			if (Arr[index] < Arr[j])
			{
				index = j;  //记录待排序元素中最大值的索引		 		
			}
		}
		if (i != index)  // 将最大值移到前面
		{
			int tmp = Arr[i];
			Arr[i] = Arr[index];
			Arr[index] = tmp;
		}
	}
     //输出排序后的元素
	for (int i = 0; i < BUFFSIZE; i++)
	{
		cout << Arr[i] << " ";
	}
	cout << endl;
}
```

### 二分查找实现（C）

```c
int binary_search(int key,int a[],int n) //自定义函数binary_search(){    
int low,high,mid,count=0,count1=0;    low=0;    high=n-1;    
while(low<high)    //査找范围不为0时执行循环体语句    
{        
 count++;          //count记录査找次数        
 mid=(low+high)/2;    //求中间位置        
 if(key<a[mid])    //key小于中间值时            
 {
 high=mid-1;
 }                 //确定左子表范围        
 else if(key>a[mid])       //key 大于中间值时            
     low=mid+1;           //确定右子表范围        
 else if(key==a[mid])    //当key等于中间值时，证明查找成功        
{            
  printf("查找成功!\n 查找 %d 次!a[%d]=%d",count,mid,key);    //输出査找次数及所査找元素在数组中的位置            
  count1++;    //count1记录查找成功次数            
  break;        
}    
}    
if(count1==0)    //判断是否查找失敗        
{
printf("查找失敗!");    //査找失敗输出no found    
}
return 0;}
```

### malloc函数

## 1.描述

* C 库函数 **void \*malloc(size_t size)** 分配所需的内存空间，并返回一个指向它的指针。

## 2.声明

* 下面是 malloc() 函数的声明。

```c
void *malloc(size_t size)
```

## 3.返回值

* 该函数返回一个指针 ，指向已分配大小的内存。如果请求失败，则返回 NULL。

## 链表中的递归（对应leetcode203）

* struct ListNode* removeElements(struct ListNode* head, int val) {
      if (head == NULL) {
          return head;
      }
      head->next = removeElements(head->next, val);
      return head->val == val ? head->next : head;
  }

struct ListNode* removeElements(struct ListNode* head, int val) {
    if (head == NULL) {
        return head;
    }
    head->next = removeElements(head->next, val);
    return head->val == val ? head->next : head;
} 

