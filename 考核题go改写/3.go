package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	var array [][]int
	for i := 0; i <= N-1; i++ {
		var arr []int
		array = append(array, arr)
	}
	var right []int
	var codes []int
	for i := 0; i <= M-1; i++ {
		var n int
		fmt.Scan(&n)
		codes = append(codes, n)
	}
	for i := 0; i <= M-1; i++ {
		var n int
		fmt.Scan(&n)
		right = append(right, n)
	}
	for i := 0; i <= N-1; i++ {
		for j := 0; j <= M-1; j++ {
			var n int
			fmt.Scan(&n)
			array[i] = append(array[i], n)
		}
	}
	for i := 0; i <= N-1; i++ {
		code := 0
		for j := 0; j <= M-1; j++ {
			if array[i][j] == right[j] {
				code = code + codes[j]
			}
		}
		fmt.Println(code)
	}
}
