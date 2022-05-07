package main

import "fmt"

// 彩云翻译返回的结果
type CaiYunDictResponse struct {
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

type YoudaoDictResponse struct {
	TranslateResult [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	ErrorCode   int    `json:"errorCode"`
	Type        string `json:"type"`
	SmartResult struct {
		Entries []string `json:"entries"`
		Type    int      `json:"type"`
	} `json:"smartResult"`
}

type QueryRequestBody struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

// 彩云翻译英文翻译成中文的请求体构造方法
func newCaiyunQuerySource(word string) QueryRequestBody {
	return QueryRequestBody{
		TransType: "en2zh",
		Source:    word,
	}
}

// 有道翻译请求表单数据构造函数
func newYoudaoQuerySource(word string) string {
	return fmt.Sprintf("i=%s&from=AUTO&to=AUTO&smartresult=dict&client=fanyideskweb&salt=16519056897718&sign=a88e85543327305152214e1e7a99da8a&lts=1651905689771&bv=f4ba6f45063bb1cd2f2872ed23e89c8f&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTlME", word)
}
