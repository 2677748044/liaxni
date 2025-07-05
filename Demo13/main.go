package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomNumber int // 全局变量，用于存储生成的随机数

func main() {
	// 设置随机种子（确保每次运行结果不同）
	rand.Seed(time.Now().UnixNano())

	// 生成0-99之间的随机整数
	randomNumber = rand.Intn(100) // 100表示上限（不包含）

	fmt.Printf("生成的随机数 (0-99): %d\n", randomNumber)

	CaiShu() // 调用猜数字函数
}

func CaiShu() {
	var num1 int
	score := 0              // 初始化分数
	startTime := time.Now() // 记录开始时间
	for {
		if time.Since(startTime).Seconds() > 60 {
			fmt.Println("时间到！你没有在60秒内猜对数字。")
			return
		}
		fmt.Printf("请输入一个你猜的数字（剩余时间：%d 秒）：", 60-int(time.Since(startTime).Seconds()))
		fmt.Scanln(&num1)
		score++
		if num1 < randomNumber {
			fmt.Println("你猜的数字小了")
		} else if num1 > randomNumber {
			fmt.Println("你猜的数字大了")
		} else {
			fmt.Printf("恭喜你，猜对了！你一共猜了 %d 次。\n", score)
			break
		}
	}
}
