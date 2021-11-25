package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/typenameman/gotools/regerep/html"
	"github.com/typenameman/gotools/regerep/jsons"
)

func main() {
	//进入初始登陆网页提交信息，获取临时token
	type Body struct {
		Username string
		Password string
	}
	var name string
	fmt.Print("请输入账号:")
	fmt.Scanln(&name)
	fmt.Print("请输入密码:")
	paswod, _ := gopass.GetPasswdMasked()
	paswodbas64 := base64.StdEncoding.EncodeToString(paswod)
	body := Body{
		Username: name,
		Password: paswodbas64,
	}
	by, _ := json.Marshal(body)
	payload := strings.NewReader(string(by))
	header := map[string]string{"Content-Type": "application/json", "Host": "pass.muxi-tech.xyz", "Content-Length": "51"}
	res := jsons.ShowJsonContents("http://pass.muxi-tech.xyz/auth/api/signin", "POST", header, payload, "token")
	token0 := res[0]

	//进入正式登录门户，传入初始token并获取正式token及cookie
	type Bodys struct {
		Email string
		Token string
	}
	bodys := Bodys{
		Email: name,
		Token: token0,
	}
	boy, _ := json.Marshal(bodys)
	boystr := string(boy)
	boystr = strings.ReplaceAll(boystr, "Email", "email")
	boystr = strings.ReplaceAll(boystr, "Token", "token")
	payload = strings.NewReader(boystr)

	Header := map[string]string{"Content-Type": "application/json"}
	res = jsons.ShowJsonContents("http://work.muxi-tech.xyz/api/v1.0/auth/login/", "POST", Header, payload, "token")
	token := res[0]
	cookie := "workbench_token=" + token

	//进入木犀工作台爬取内容，现随机选取其中一个网页进行爬取信息
	Header["token"] = token
	Header["Cookie"] = cookie
	Header["Host"] = "work.muxi-tech.xyz"
	returnres := jsons.ShowJsonContentsWithUnicode("http://work.muxi-tech.xyz/api/v1.0/file/doc/1155/", "GET", Header, nil, "content")
	for _, cont := range returnres {
		resu := html.ShowTags(cont)
		fmt.Println(resu)
	}
	returnres = jsons.ShowJsonContentsWithUnicode("http://work.muxi-tech.xyz/api/v1.0/file/doc/1155/", "GET", Header, nil, "create_time")
	fmt.Println(returnres)
	returnres = jsons.ShowJsonContentsWithUnicode("http://work.muxi-tech.xyz/api/v1.0/file/doc/1155/", "GET", Header, nil, "creator")
	fmt.Println(returnres)
	returnres = jsons.ShowJsonContentsWithUnicode("http://work.muxi-tech.xyz/api/v1.0/file/doc/1155/", "GET", Header, nil, "name")
	fmt.Println(returnres)
}
