package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
func queryFromCaiyun(word string, ch chan []string) {
	client := &http.Client{}

	// 修改的代码
	// var data = strings.NewReader(`{"trans_type":"en2zh","source":"goo"}`)
	queryBody, err := json.Marshal(newCaiyunQuerySource(word))
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
	var dictResponse CaiYunDictResponse
	json.Unmarshal(bodyBytes, &dictResponse)

	var msg []string
	msg = append(msg, fmt.Sprint("UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs))
	for _, item := range dictResponse.Dictionary.Explanations {
		// fmt.Println(item)
		msg = append(msg, item)
	}
	ch <- msg
}

func queryFromYoudao(word string, ch chan []string) {
	client := &http.Client{}
	var data = strings.NewReader(newYoudaoQuerySource(word))
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "OUTFOX_SEARCH_USER_ID=50780092@10.108.162.134; JSESSIONID=aaaNpYcG7CbLvcNFfqDcy; OUTFOX_SEARCH_USER_ID_NCOO=2052774788.767026; ___rl__test__cookies=1651905689766")
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
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
	var dictResponse YoudaoDictResponse
	json.Unmarshal(bodyBytes, &dictResponse)
	var msg []string

	for _, item := range dictResponse.SmartResult.Entries {
		msg = append(msg, item)
	}

	ch <- msg
}
