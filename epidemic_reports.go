package main

import (
	"encoding/json"
	"epidemic_reports/config"
	"fmt"
	"net/http"
	"strings"
)

// URL 请求路径
const URL = "https://appportalserver.g5air.com/YQempInfo/PostEdit"

func setCount(user config.User) error {
	reqCountBytes, err := json.Marshal(user) // 把请求结构体解析为json
	if err != nil {
		fmt.Println("marshal failed. the error info: ", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(string(reqCountBytes))) // 因为要调用的Rest接口是PUT类型的，需要先构造Request
	if err != nil {
		fmt.Println("new request failed with error: %s", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json") // 注意要给Request的请求头添加上内容类型
	client := http.Client{}                            // 创建一个httpClient
	_, err = client.Do(req)                            // 调用rest接口
	if err != nil {
		return err
	}
	return nil
}
