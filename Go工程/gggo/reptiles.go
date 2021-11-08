package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	me := "[\u4e00-\u9fa5]{3}(\\s)*[0-9]{18}"
	method, _ := regexp.Compile(me)
	res, _ := http.Get("https://www.wangsu123.cn/shenfenzheng/hangzhou/")
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	content := string(body)
	result := method.FindAllString(content, -1)
	for _, b := range result {
		fmt.Println(b)
		fmt.Println("\n")
	}

}
