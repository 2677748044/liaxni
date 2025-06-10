package main

import (
	"demo2/test"
	"fmt" //Go 1.17 及以后默认启用 Go Modules（go.mod），即使你设置了 GOPATH，也可能会强制用模块模式。你可以用下面命令查看当前是否启用模块：
)

func main() {
	fmt.Println(test.SnoTo) //test包的变量
}
