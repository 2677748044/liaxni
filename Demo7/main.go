package main

import "fmt"

func main() {
	/*	var sum int = 0
		for i := 1; i <= 200; i++ {
			sum += i
			fmt.Println(sum)
			if sum >= 200 {
				break
			}
		}*/
lable: //标签
	for X := 1; X <= 3; X++ {
		for Y := 2; Y <= 5; Y++ {
			fmt.Printf("X: %v,Y: %v \n", X, Y) //双for循环
			if X == 2 && Y == 2 {
				break lable //停止的是 x=2 && y =2 以后的输出
			}
		}
	}
}
