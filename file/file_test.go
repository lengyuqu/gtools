package file

import (
	"testing"
)

func TestReadAll(t *testing.T) {
	data, err := ReadAll("file.go")
	if err != nil {
		t.Errorf("module file function ReadAll is error, reason:%s", err.Error())
		return
	}
	println(string(data))
}

func TestFile(t *testing.T) {

	notExsitFile := "D://1jdh134jfd/1923834dfda/21321.jpd"

	if IsExsit(notExsitFile) {
		t.Error("file.Exsit test error")
	}

	if !IsFile("file.go") {
		t.Error("file.IsFile test error")
	}

	if !IsDir("../file") {
		t.Error("file.IsDir test error")
	}

	if GetName(notExsitFile) != "21321.jpd" {
		t.Error("file.GetName test error")
	}

	if GetExt(notExsitFile) != ".jpd" {
		t.Error("file.GetExt test error")
	}

	if GetNameWithoutExt(notExsitFile) != "21321" {
		t.Error("file.GetNameWithoutExt test error")
	}

	dir, f := SplitDirAndFileName(notExsitFile)
	t.Log("dir:", dir, "filename:", f)

	if err := Create("E:/golang/src/temp.go"); err != nil {
		t.Error("file.Create test error")
	}
}

func TestZip(t *testing.T) {
	if err := Zip("D:/me/myself/gosrc/gweb/static/report", "test"); err != nil {
		t.Error("file.Zip test error:", err.Error())
	}

	if err := Unzip("D:/me/myself/gosrc/gweb/static/report/test.zip", "D:/me/myself/gosrc/gweb/static/report/test"); err != nil {
		t.Error("file.Unzip test error:", err.Error())
	}
}
