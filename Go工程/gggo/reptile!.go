package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	var urls []string
	var result [][]string
	url1 := "https://www.baidu.com/link?url=MFsIIpOC0zsMo6N0eNZHPBfP-MxEe1x9ZwBACiqloivl8Ir-4FflOxhV1GLcldgB938XvhT9dwxqpHQz7Ld91nk_3_H7ofebUhL7INO5iuG&wd=&eqid=cd920071000347d900000002617e045b"
	url2 := "https://henan.qq.com/a/20171107/069413.htm"
	url3 := "http://liaocheng.dzwww.com/lcxw/202009/t20200915_6607808.htm"
	url4 := "http://henan.sina.com.cn/news/s/2018-04-25/detail-ifzqvvsc1649660.shtml"
	url5 := "https://m.thepaper.cn/baijiahao_14747519"
	url6 := "https://m.thepaper.cn/baijiahao_12342497"
	urls = append(urls, url1)
	urls = append(urls, url2)
	urls = append(urls, url3)
	urls = append(urls, url4)
	urls = append(urls, url5)
	urls = append(urls, url6)
	for _, url := range urls {
		res, _ := http.Get(url)
		body, _ := ioutil.ReadAll(res.Body)
		content := string(body)
		rege := "[0-9]{18}"
		pattern, _ := regexp.Compile(rege)
		re := pattern.FindAllString(content, -1)
		result = append(result, re)
	}
	f, _ := os.Create("./result.txt")
	defer f.Close()
	for _, sli := range result {
		for _, slice := range sli {
			f.WriteString(slice)
			f.WriteString("\n")
		}
	}

}
