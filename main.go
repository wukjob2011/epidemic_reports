package main

import (
	"bytes"
	"encoding/json"
	"epidemic_reports/config"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	ConfigPath = "config.json"
	URL        = "https://appportalserver.g5air.com/YQempInfo/PostEdit"
)

func main() {
	log.Println("Starting...")
	conf := initConfig()
	fmt.Println(conf)
	// 新建一个定时任务对象
	c := cron.New()
	_ = c.AddFunc("0 30 9 0 0 ?", func() {
		// 给对象增加定时任务
		report(conf)
	})
	c.Start()
	select {}
}

/**
初始化配置
*/
func initConfig() config.Config {
	formBytes := loadFileForm(ConfigPath)
	return unMarshalConfig(formBytes)
}

/**
反序列化配置文件
*/
func unMarshalConfig(bytes []byte) config.Config {
	conf := config.Config{}
	if err := json.Unmarshal(bytes, &conf); err != nil {
		log.Println("配置解析错误", err.Error())
		os.Exit(-1)
	}
	return conf
}

/**
上报疫情健康信息
*/
func report(config config.Config) {
	for _, report := range config.Reports {
		sleepTime := generateSleepTime()
		time.Sleep(time.Duration(sleepTime) * time.Second)
		postData := buildFromData(report)
		go PostWithFormData(http.MethodPost, URL, &postData)
	}
}

/**
生成睡眠时间
*/
func generateSleepTime() int {
	return rand.Intn(1800)
}

/**
构建表单请求
*/
func buildFromData(report config.Report) map[string]string {
	formData := make(map[string]string)
	formData["Account"] = report.Account
	formData["DeptCode"] = report.DeptCode
	formData["CompanyName"] = report.CompanyName
	formData["DeptName"] = report.DeptName
	formData["UserName"] = report.UserName
	formData["Sex"] = report.Sex
	formData["Age"] = report.Age
	formData["Gwmc"] = report.Gwmc
	formData["Phone"] = report.Phone
	formData["WorkPlace"] = report.WorkPlace
	formData["UpTime"] = report.UpTime
	formData["Province"] = report.Province
	formData["City"] = report.City
	formData["Region"] = report.Region
	formData["WzStatus"] = report.WzStatus
	formData["Status"] = report.Status
	formData["CommunityRec"] = report.CommunityRecord
	formData["FamilyStatus"] = report.FamilyStatus
	formData["IsToCompany"] = report.IsToCompany
	formData["Wz_remark"] = report.WzRemark
	formData["Temperature"] = report.Temperature
	return formData
}

// PostWithFormData 健康信息上报
func PostWithFormData(method, url string, postData *map[string]string) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		_ = w.WriteField(k, v)
	}
	_ = w.Close()
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, _ := http.DefaultClient.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Printf("用户%s: 健康信息上报%s", (*postData)["Account"], data)
}

/**
加载文件
*/
func loadFileForm(path string) []byte {
	formBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("读取配置文件出错!!!", err.Error())
		os.Exit(-1)
	}
	return formBytes
}
