package main

import "fmt"

func cal(num1, num2 int) (int, int) { //自定义函数 提高复用率
	var num int = 0
	num += num1
	num += num2

	var result int = num1 - num2
	return num, result
}
func main() {
	num1, result := cal(10, 50)
	fmt.Println(num1)
	fmt.Println(result)
}
