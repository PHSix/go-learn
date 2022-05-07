package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
	// 产生一个随机数
	securtNumber := rand.Intn(100)
	// 输出游戏规则
	fmt.Println("This is a guessing game:")
	fmt.Println("The rules of this game is:")
	fmt.Println("\t1.We will get a random number at first(between 0 to 100)")
	fmt.Println("\t2.User input a number")
	fmt.Println("\t3.Output the result of contrast(if smaller then back to the first setup, else break)")
	// 构建一个读取用户输入的一个stream
	render := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input your guess number: ")
		input, err := render.ReadString('\n')
		if err != nil {
			fmt.Println("Read user input error.")
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		inputNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invaild Input")
			continue
		}
		if inputNum >= securtNumber {
			fmt.Printf("Yes, number %d is bigger than %d\n", inputNum, securtNumber)
			break
		} else {
			fmt.Printf("Number %d is smaller than or equal to securt number, please have a guess again.\n", inputNum)
			continue
		}
	}
}
