使用 [fmt.Sprintf](https://link.juejin.cn?target=https%3A%2F%2Fgolang.org%2Fpkg%2Ffmt%2F%23Sprintf) 将[interface value](https://link.juejin.cn?target=https%3A%2F%2Fyourbasic.org%2Fgolang%2Finterfaces-explained%2F)转换为字符串。

```go
var x interface{} = "abc"
str := fmt.Sprintf("%v", x)
复制代码
```

实际上，可以使用相同的技术来获取任何数据结构的字符串表示形式。

```go
var x interface{} = []int{1, 2, 3}
str := fmt.Sprintf("%v", x)
fmt.Println(str) // "[1 2 3]"
```

