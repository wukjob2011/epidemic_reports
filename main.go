package main

import (
	"encoding/json"
	"epidemic_reports/config"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const ConfigPath = "config.json"

func main() {
	config := initConfig()
	fmt.Println(config)

}

func initConfig() config.Config {
	bytes := loadFileForm(ConfigPath)
	return unMarshalConfig(bytes)
}

func unMarshalConfig(bytes []byte) config.Config {
	config := config.Config{}
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Println("配置解析错误", err.Error())
		os.Exit(-1)
	}
	return config
}

/**
加载文件
*/
func loadFileForm(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("读取配置文件出错!!!", err.Error())
		os.Exit(-1)
	}
	return bytes
}
