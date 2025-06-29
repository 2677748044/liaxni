package main

import (
	"Demo12/tesutils"
	_ "Demo12/tesutils"
	"fmt"
)

var num int = test() //测试 全剧变量的初始化顺序

func test() int {
	fmt.Println("test函数执行")
	return 10
}

func init() {
	fmt.Println("init初始化函数执行")
}

func main() {
	fmt.Println("main函数被执行")
	fmt.Println("age= ", tesutils.Age, "sex= ", tesutils.Sex, "name= ", tesutils.Name)
}
