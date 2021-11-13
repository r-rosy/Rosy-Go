package main

import "fmt"

func check(arr [][]int, a int, b int) bool {
	for i := 0; i <= a-1; i++ {
		for j := 0; j <= b-1; j++ {
			if arr[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
func change(arr [][]int, a int, b int) {
	arrs := make([][]int, a)
	for i, _ := range arrs {
		arrs[i] = make([]int, b)
		copy(arrs[i], arr[i]) //copy函数必须提前制定切片长度
	}
	for i := 0; i <= a-1; i++ {
		for j := 0; j <= b-1; j++ {
			if arr[i][j] == 1 {
				if j == 0 {
					arrs[i][j+1] = 1
				}
				if j == b-1 {
					arrs[i][j-1] = 1
				}
				if j != 0 && j != b-1 {
					arrs[i][j+1] = 1
					arrs[i][j-1] = 1
				}
				if i == 0 {
					arrs[i+1][j] = 1
				}
				if i == a-1 {
					arrs[i-1][j] = 1
				}
				if i != 0 && i != a-1 {
					arrs[i+1][j] = 1
					arrs[i-1][j] = 1
				}
			}
		}
	}
	for i, _ := range arr {
		copy(arr[i], arrs[i])
	}
	/*此处不能用切片直接赋值*/
}
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var arr [][]int
	for i := 0; i <= a-1; i++ {
		var array []int
		for j := 0; j <= b-1; j++ {
			var n int
			fmt.Scan(&n)
			array = append(array, n)
		}
		arr = append(arr, array)
	}

	n := 0
	for {
		change(arr, a, b)
		n++
		if check(arr, a, b) == true {
			break
		}
	}
	fmt.Println(n)
}
