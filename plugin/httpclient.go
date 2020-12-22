package plugin

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 基础方法，这里多用于访问webapi，配合上json转换。此方法可以运行但是不算完善。
func httpDo(method string, url string, msg string) (int, string, error) {
	// log.Println("----", url, "----")
	client := &http.Client{}
	body := bytes.NewBuffer([]byte(msg))
	req, err := http.NewRequest(method,
		url,
		body)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	defer resp.Body.Close()

	result_code := resp.StatusCode

	result_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	return result_code, string(result_body), nil

}

// post方式
func HttpDoPost(url string, msg string) (int, string, error) {
	res_code, res_str, err := httpDo("POST", url, msg)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	return res_code, res_str, nil
}

// put方式
func HttpDoPut(url string, msg string) (int, string, error) {
	res_code, res_str, err := httpDo("PUT", url, msg)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	return res_code, res_str, nil
}

// patch方式
func HttpDoPatch(url string, msg string) (int, string, error) {
	res_code, res_str, err := httpDo("PATCH", url, msg)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	return res_code, res_str, nil
}

// get方式
func HttpDoGet(url string, msg string) (int, string, error) {
	res_code, res_str, err := httpDo("GET", url, msg)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	return res_code, res_str, nil
}

// delete方式
func HttpDoDelete(url string, msg string) (int, string, error) {
	res_code, res_str, err := httpDo("DELETE", url, msg)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	return res_code, res_str, nil
}

func HttpDo1(method string, url string, msg string) (int, string, error) {
	// log.Println("----", url, "----")
	client := &http.Client{}
	body := strings.NewReader(msg)
	// body := bytes.NewBuffer([]byte(msg))
	req, err := http.NewRequest(method,
		url,
		body)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	defer resp.Body.Close()

	result_code := resp.StatusCode

	result_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	return result_code, string(result_body), nil

}
