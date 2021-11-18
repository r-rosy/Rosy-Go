package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anaskhan96/soup"
)

func main() {
	req := "https://news.baidu.com/"
	res, _ := http.Get(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	root := soup.HTMLParse(string(body))
	//fmt.Println(string(body))
	rot := root.FindAll("div", "class", "hotnews")
	for _, b := range rot {
		content1 := b.FindAll("a")
		for _, b := range content1 {
			fmt.Println(b.Text())
		}
		content2 := b.FindAll("b")
		for _, b := range content2 {
			fmt.Println(b.Text())
		}

	}
	/*rot = root.FindAll("ul", "class", "ulist focuslistnews")
	for _, b := range rot {
		content1 := b.FindAll("a")
		for _, b := range content1 {
			fmt.Println(b.Text())
		}
		content2 := b.FindAll("b")
		for _, b := range content2 {
			fmt.Println(b.Text())
		}

	}*/
}
