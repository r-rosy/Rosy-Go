package regerep

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func ShowTags(url string) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	reg, _ := regexp.Compile("<.{1,}?>[\u4e00-\u9fa5||0-9||a-z||A-Z||\u0020||\u3002||\uff1b||\uff0c||\uff1a||\u201c||\u201d||\uff08||\uff09||\u3001||\uff1f||\u300a||\u300b||\u007c||\uff01||\u003a||\uff5c||\u002c\u0022\u003f\uff5e\u2103\u0025\u002b]{1,}?<.{1,}?>")
	pat, _ := regexp.Compile(">[\u4e00-\u9fa5||0-9||a-z||A-Z||\u0020||\u3002||\uff1b||\uff0c||\uff1a||\u201c||\u201d||\uff08||\uff09||\u3001||\uff1f||\u300a||\u300b||\u007c||\uff01||\u003a||\uff5c||\u002c\u0022\u003f\uff5e\u2103\u0025\u002b]{1,}?<")
	y := reg.FindAllStringSubmatch(string(body), -1)
	for _, b := range y {
		for _, a := range b {
			str := pat.FindAllStringSubmatch(a, -1)
			for _, p := range str {
				for i, _ := range p {
					p[i] = strings.ReplaceAll(p[i], ">", "")
					p[i] = strings.ReplaceAll(p[i], "<", "")
					fmt.Println(p[i])
				}
			}
		}
	}
}
