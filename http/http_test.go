package http

import (
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://www.baidu.com"

	data, err := Get(url, nil, nil)
	if err != nil {
		t.Logf("request to baidu error %s", err.Error())
	}

	t.Log(string(data))
}

func TestPostJSON(t *testing.T) {
	url := "http://localhost:8088/file"
	var paramJson strings.Builder
	paramJson.WriteString("{")
	paramJson.WriteString("\"id\": \"1347283170209112578\",")
	paramJson.WriteString("\"fileName\": \"ctest\"")
	paramJson.WriteString("}")

	data, err := PostJSON(url, paramJson.String())
	if err != nil {
		t.Logf("request to baidu error %s", err.Error())
	}

	t.Log(string(data))
}

func TestPostForm(t *testing.T) {
	url := "http://localhost:8088/file"

	params := map[string]string{
		"id":       "1347283170209112578",
		"fileName": "ctest",
	}

	data, err := PostForm(url, params)
	if err != nil {
		t.Logf("request to baidu error %s", err.Error())
	}

	t.Log(string(data))
}

func TestPostFiles(t *testing.T) {
	url := "http://localhost:8088/file/audio"
	params := map[string]string{
		"serialNo": "1000770",
	}

	files := map[string]string{
		"file": "D:\\语音数据集\\BAC009S0025W0392.wav",
	}
	data, err := PostFiles(url, params, files)
	if err != nil {
		t.Logf("request to baidu error %s", err.Error())
	}

	t.Log(string(data))
}

func TestPostFile(t *testing.T) {
	url := "http://localhost:8088/file/audios"
	params := map[string]string{
		"serialNo": "1000770",
	}

	data, err := PostFile(url, params, "file", "D:\\语音数据集\\BAC009S0025W0392.wav")
	if err != nil {
		t.Logf("request to baidu error %s", err.Error())
	}

	t.Log(string(data))
}
