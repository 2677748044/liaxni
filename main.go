package main

import "fmt"

func main() {
	var X int = 1
	var cout int
	fmt.Printf("请输入：")
	fmt.Scanln(&cout)

	switch cout / 10 { //switch 后面是一个表达式，表达式的值会与每个case进行比较
	case 9: // cout / 10 计算cout的十位数
		fmt.Println("等级是SS")
	case 8:
		fmt.Println("等级为A")
	case 7:
		fmt.Println("等级为B")
	case 6, 5, 4, 3, 2, 1, 0:
		fmt.Println("等级为C")
	}

	switch { // switch 后面没有表达式，直接进行条件判断 可当if
	case X == 1:
		fmt.Println("aaa")
	}

	switch b := 22; { //switch 后面定义表达式 需要分号结尾
	case b > 10:
		fmt.Println("b大于10")
	case b < 10:
		fmt.Println("b小于10")
	}

}
