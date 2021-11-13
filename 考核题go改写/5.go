package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MySprintf(s *string, str string, args []int) {
	index := strings.Index(str, "%d")
	if index == -1 {
		*s = str
	} else {
		for _, i := range args {
			is := strconv.Itoa(i)
			if strings.Index(str, "%d") == -1 {
				break
			}
			str = strings.Replace(str, "%d", is, 1)
			str = strings.Replace(str, "%%", "%", 1)
		}
	}
	*s = str
}
func main() {
	var string string
	s := &string //go中string是一种值传递类型
	var args []int
	args = append(args, 0)
	args = append(args, 1)
	args = append(args, 2)
	MySprintf(s, "ss%dg%%gg%d", args)
	fmt.Println(*s)
}
