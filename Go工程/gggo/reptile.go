package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req := "https://www.liaoxuefeng.com/"
	res, _ := http.Get(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
