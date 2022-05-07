/*
程序说明：
	1. 运行程序并在后面追加上一个单词
	2. 程序发送一次http请求单彩云科技的翻译api上
	3. 获取翻译内容返回并输出

API地址：https://api.interpreter.caiyunai.com/v1/dict
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 翻译的Api
const transUrl = "https://api.interpreter.caiyunai.com/v1/dict"

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []interface{} `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}
type QueryRequestBody struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

// 英文翻译成中文的请求体构造方法
func newQueryFromEn2Zh(word string) QueryRequestBody {
	return QueryRequestBody{
		TransType: "en2zh",
		Source:    word,
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Invaild input.")
	}
	queryEn2Zh(args[1])
}

/*
代码自动生成自：https://curlconverter.com/#go
原curl命令(第二行有添加了一个转义字符)：
curl 'https://api.interpreter.caiyunai.com/v1/dict' \
  -H 'Accept: application/json, text/plain, *\/\*' \
  -H 'Accept-Language: zh-CN,zh;q=0.9' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json;charset=UTF-8' \
  -H 'Origin: https://fanyi.caiyunapp.com' \
  -H 'Referer: https://fanyi.caiyunapp.com/' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Site: cross-site' \
  -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36' \
  -H 'X-Authorization: token:qgemv4jr1y38jyq6vhvi' \
  -H 'app-name: xy' \
  -H 'device-id: ' \
  -H 'os-type: web' \
  -H 'os-version: ' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  --data-raw '{"trans_type":"en2zh","source":"goo"}' \
  --compressed
*/
func queryEn2Zh(word string) {
	client := &http.Client{}

	// 修改的代码
	// var data = strings.NewReader(`{"trans_type":"en2zh","source":"goo"}`)
	queryBody, err := json.Marshal(newQueryFromEn2Zh(word))
	if err != nil {
		log.Fatal("Json marshal failed.", err)
	}
	data := bytes.NewReader(queryBody)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}

	// 修改完成

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dictResponse DictResponse
	json.Unmarshal(bodyBytes, &dictResponse)

	fmt.Printf("Source Input: %s\n", word)
	fmt.Println("UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}
