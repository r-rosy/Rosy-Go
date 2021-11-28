# Go语言接口类型转换教程

在 **[Golang](https://haicoder.net/golang/golang-tutorial.html)** 中，将一个 **[接口](https://haicoder.net/golang/golang-interface.html)** 类型转换成另一个接口 **[类型](https://haicoder.net/golang/golang-datatype.html)**，或者将一个接口转换为另一个基本类型，都必须需要使用 **[类型断言](https://haicoder.net/golang/golang-type-assert.html)**。

## Go语言接口类型转换

### 语法

```
value, ok := x.(T)
```

### 参数

| 参数 | 描述             |
| ---- | ---------------- |
| *x*  | 需要转换的数据。 |
| *T*  | 需要转换的类型。 |

### 返回值

| 返回值  | 描述                                                         |
| ------- | ------------------------------------------------------------ |
| *value* | 转换后的数据。                                               |
| *ok*    | 转换成功与否的 **[bool](https://haicoder.net/golang/golang-bool.html)** 变量表示。 |

### 说明

将接口 x 转换成 T 类型。 如果转换成功，返回转换成功后的值，即 value，ok 为 **[true](https://haicoder.net/golang/golang-bool.html)**。如果转换失败，value 为 零值，ok 为 false。

## Go语言接口类型转换

### 语法

```
value := x.(T)
```

### 说明

将接口 x 转换成 T 类型。 如果转换成功，返回转换成功后的值，即 value，如果转换失败，程序会 **[panic](https://haicoder.net/golang/golang-panic.html)**。

## 案例

### 接口类型转换

使用接口类型断言，将接口类型转成 **[string](https://haicoder.net/golang/golang-string.html)**

```
package main
import (
	"fmt"
)
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	// 使用接口类型转换，将接口类型转成 string
	var svalue interface{}
	svalue = "HaiCoder"
	if value, ok := svalue.(string); ok{
		fmt.Println("Ok Value =", value, "Ok =", ok)
	}else{
		fmt.Println("Failed Value =", value, "Ok =", ok)
	}
}
```

程序运行后，控制台输出如下：

![13_golang类型断言.png](https://haicoder.net/uploads/pic/server/golang/golang-interface/13_golang%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80.png)

首先，我们定义了一个接口类型的 **[变量](https://haicoder.net/golang/golang-variable.html)** svalue，并给其赋值为 “HaiCoder”，接着，我们使用接口类型转换将变量 “HaiCoder” 转成 string 类型。

接口类型转换转换，返回两个值，一个表示转换后的值，一个 bool 类型的变量表示是否转换成功。最后，我们发现，成功将接口类型转成了 string 类型，并返回了转换成功后的值。

### 接口类型转换

使用接口类型断言，将接口类型转成 **[指针类型](https://haicoder.net/golang/golang-pointer.html)**

```
package main
import (
	"fmt"
)
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	// 使用接口类型转换，将接口类型转成指针类型
	var svalue interface{}
	str := "HaiCoder"
	svalue = &str
	if value, ok := svalue.(*string); ok{
		fmt.Println("Ok Value =", *value, "Ok =", ok)
	}else{
		fmt.Println("Failed Value =", value, "Ok =", ok)
	}
}
```

程序运行后，控制台输出如下：

![14_golang类型断言.png](https://haicoder.net/uploads/pic/server/golang/golang-interface/14_golang%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80.png)

首先，我们定义了一个接口类型的变量 svalue，接着，我们定义了一个字符串类型的变量 str，并且将字符串变量的地址赋值给接口。

因此，此时接口的类型为字符串指针，所以，我们使用类型断言将接口转成了字符串指针类型。

### 接口类型转换失败

使用接口类型转换，如果断言失败，不使用 bool 变量接受，程序 panic

```
package main
import (
	"fmt"
)
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	// 使用接口类型转换，如果断言失败，不使用 bool 变量接受，程序 panic
	var intvalue interface{}
	intvalue = 100
	value := intvalue.(float32);
	fmt.Println("Ok Value =", value)
}
```

程序运行后，控制台输出如下：

![15_golang类型断言.png](https://haicoder.net/uploads/pic/server/golang/golang-interface/15_golang%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80.png)

在使用接口类型转换时，如果我们只使用一个值接受断言的返回值，那么此时如果断言失败，那么程序会 panic。

## Golang接口类型转换总结

在 Golang 中，将一个接口类型转换成另一个接口类型，或者将一个接口转换为另一个基本类型，都必须需要使用类型断言。Go 语言接口类型转换语法：

```
value, ok := x.(T)
```

将接口 x 转换成 T 类型。 如果转换成功，返回转换成功后的值，即 value，ok 为 true。如果转换失败，value 为 零值，ok 为 false。Go 语言接口类型转换语法：

```
value := x.(T)
```

将接口 x 转换成 T 类型。 如果转换成功，返回转换成功后的值，即 value，如果转换失败，程序会 panic。