package main

import (
	"fmt"
)

func main() {
	var n1 int = 10
	fmt.Println(n1)

	var ptr *int = &n1 //ptr 执政名 *int 类型 &符号+变量 获取地址   * 获取地址里面的值
	//var ptr *int =  n1    指针变量接受的的一定是一个地址值 &
	// var ptr *float32 = &n1   代表指针指向的是float32类型的 指针变量的地址不能不匹配 n1 是int
	*ptr = 22
	fmt.Println(n1)

}
