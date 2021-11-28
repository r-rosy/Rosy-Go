package database

const (
	Username = "root"
	Password = "123456"
	Url      = "root:123456@tcp(127.0.0.1:3306)/user"
)

var Types1 = map[string]string{"name": "char(10)", "pass": "char(20)", "niname": "char(10)"}
var Types2 = map[string]string{"name": "char(10)", "value": "char(20)", "expire": "char(100)"}
