package http

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Get 请求
//
// link 请求地址 http://www.baidu.com
//
// header 请求头
//
// params 请求参数
func Get(link string, header map[string]string, params map[string]string) ([]byte, error) {

	client := &http.Client{Timeout: 5 * time.Second}

	//忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	p := url.Values{}
	u, _ := url.Parse(link)

	// 设置url参数
	for k, v := range params {
		p.Set(k, v)
	}

	u.RawQuery = p.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	// 设置header
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http is exception")
	}

	return ioutil.ReadAll(resp.Body)
}

// PostJSON post请求提交JSON数据
//
// link 请求地址 http://www.baidu.com
//
// body 提交的json数据
func PostJSON(link string, body string) ([]byte, error) {

	resp, err := http.Post(link, "application/json; charset=utf-8", strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// PostForm post请求提交表单
//
// link 请求URL
//
// values 表单参数
func PostForm(link string, values map[string]string) ([]byte, error) {

	p := url.Values{}

	for k, v := range values {
		p.Set(k, v)
	}

	resp, err := http.PostForm(link, p)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// PostFiles post表单上传文件
//
// params 表单参数
//
// files [文件参数名]文件存放路径
func PostFiles(link string, params map[string]string, files map[string]string) ([]byte, error) {

	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for name, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		part, err := writer.CreateFormFile(name, path)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", link, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// PostFiles 上传单个文件
//
// link 上传URL地址
//
// params 上传的表单参数
//
// filename 上传文件的参数名，比如 file
//
// path 上传文件路径，如： /home/123.wav
func PostFile(link string, params map[string]string, filename string, path string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	part, err := writer.CreateFormFile(filename, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", link, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
