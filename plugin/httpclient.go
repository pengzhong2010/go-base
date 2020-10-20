package plugin

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// 基础方法，这里多用于访问webapi，配合上json转换。此方法可以运行但是不算完善。
func httpDo(method string, url string, msg string) (string, error) {
	// log.Println("----", url, "----")
	client := &http.Client{}
	body := bytes.NewBuffer([]byte(msg))
	req, err := http.NewRequest(method,
		url,
		body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	result_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(result_body), nil

}

// post方式
func HttpDoPost(url string, msg string) {
	httpDo("POST", url, msg)
}

// get方式
func HttpDoGet(url string, msg string) (string, error) {
	res_str, err := httpDo("GET", url, msg)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return res_str, nil
}
