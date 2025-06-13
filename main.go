package main

import (
	"fmt"
)

func main() {
	var cout int64
	fmt.Printf("请输入库存剩余：") //结合 1.用户输入 2.进行条件判断（双分支）
	fmt.Scanln(&cout)
	if cout < 50 {
		fmt.Println("对不起,你输入的值小于50 库存不足")
	} else if cout < 60 {
		fmt.Println("你输入的值小于60 并且大于50")
	} else if cout < 1214 {
		fmt.Println("输入值小于1214 并且大于60 库存非常充足")
	} else if cout == 12345 {
		fmt.Println("输入的值是指定值 到达临界值")
	}
}
