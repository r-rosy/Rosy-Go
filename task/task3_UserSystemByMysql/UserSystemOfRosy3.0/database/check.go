package database

import (
	"net/http"

	"github.com/typenameman/gotools/handle"
)

func CheckCookie(r *http.Request) bool {
	tab := SqlInitCookie()

	cookies := r.Cookies()
	for _, v := range cookies {
		res := tab.CheckDetails(v.Name, "name", "name")
		if len(res) != 0 {
			res = tab.CheckDetails(v.Name, "name", "value")
			val := handle.InterfaceToString(res[0]["value"])
			if val == v.Value {
				return true
			}

		}
	}
	return false

}
