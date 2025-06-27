package main

import "fmt"

func main() {
	var x string
	fmt.Printf("请输入：")
	fmt.Scanln(&x)
	for i, value := range x { //range 遍历
		fmt.Printf("索引为 %d 具体的值为 %c \n", i, value)
	}
}
