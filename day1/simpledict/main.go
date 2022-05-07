/*
程序说明：
	1. 运行程序并在后面追加上一个单词
	2. 程序发送一次http请求单彩云科技的翻译api上
	3. 获取翻译内容返回并输出
彩云API地址：https://api.interpreter.caiyunai.com/v1/dict
有道API地址：https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// 翻译的Api
const caiyunTransUrl = "https://api.interpreter.caiyunai.com/v1/dict"
const youdaoTransurl = "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule"

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Invaild input.")
	}
	var caiyun = make(chan []string)
	var youdao = make(chan []string)
	go queryFromCaiyun(args[1], caiyun)
	go queryFromYoudao(args[1], youdao)

	caiyunResult := <-caiyun
	youdaoResult := <-youdao
	fmt.Println("彩云翻译结果\n\t", strings.Join(caiyunResult, "\n\t"))
	fmt.Println("有道翻译结果\n\t", strings.Join(youdaoResult, "\t"))
}
