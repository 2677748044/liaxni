package main

import "fmt"

func main() {
	var n1 int = +10
	fmt.Println(n1)
	var n2 int = 10 + 12
	fmt.Println(n2)
	var n3 string = "abc" + "def"
	fmt.Println(n3)

	fmt.Println(10 / 3)   //两个整数相除，结果为整数
	fmt.Println(10.0 / 3) //一个整数和一个浮点数相除，结果为浮点数
	// % 取模 a % b 等价于 a - a / b * b
	fmt.Println(10 % 3) // 10-3= 10-10/3*3
	//++ 自增操作
	var n4 int = 10
	var n5 int = n4 + 1
	n4++ //++n4   go 语言不支持前置自增
	// n5 = 11
	// n5 +=10   等价于 n5 = n 5+10
	fmt.Println(n5)
	fmt.Println(n4)

	var a int = 2
	var b int = 3
	var t int
	t = a //互换
	a = b
	b = t
	fmt.Println(a, b)
}
