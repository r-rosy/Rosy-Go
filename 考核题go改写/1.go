package main

import "fmt"

func main() {
	var h, m, s, addSec int
	fmt.Println("请输入小时:")
	fmt.Scanln(&h)
	fmt.Println("请输入分钟:")
	fmt.Scanln(&m)
	fmt.Println("请输入秒数:")
	fmt.Scanln(&s)
	fmt.Println("请输入增加的秒数:")
	fmt.Scanln(&addSec)
	s = s + addSec
	for s >= 60 {
		s = s - 60
		m++
	}
	for m >= 60 {
		m = m - 60
		h++
	}
	for h >= 24 {
		h = h - 24
	}
	fmt.Printf("%d %d %d", h, m, s)
}
