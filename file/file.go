package file

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ReadAll 读取文件所有内容
func ReadAll(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

// IsExsit 文件或目录是否存在
func IsExsit(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

// IsFile 判断一个路径是否为文件
func IsFile(path string) bool {
	fi, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// IsDir 判断一个路径是否为目录
func IsDir(path string) bool {
	fi, e := os.Stat(path)
	if e != nil {
		return false
	}
	return fi.IsDir()
}

// GetName 获取文件名
func GetName(filePath string) string {
	return path.Base(filePath)
}

// GetExt 获取文件扩展名
func GetExt(filePath string) string {
	return path.Ext(filePath)
}

// GetNameWithoutExt 获取不带扩展名的文件名
func GetNameWithoutExt(filePath string) string {
	full := path.Base(filePath)
	suffix := path.Ext(filePath)
	return full[0 : len(full)-len(suffix)]
}

// SplitDirAndFileName 将目录与文件名分开
func SplitDirAndFileName(filePath string) (dir, file string) {
	return path.Split(strings.ReplaceAll(filePath, "\\", "/"))
}

// Create 创建文件
func Create(filePath string) error {
	if !IsExsit(filePath) {
		dPath, fPath := SplitDirAndFileName(filePath)
		if dPath != "" && !IsExsit(dPath) {
			if err := os.MkdirAll(dPath, os.ModePerm); err != nil {
				return err
			}
		}

		if fPath != "" {
			file, err := os.Create(filePath)
			defer file.Close()
			if err != nil {
				return err
			}
		}
	}

	return nil

}

// Zip zip压缩，生成的压缩文件存在于压缩目录中
//
// filePath 被压缩文件目录 eg: /opt/test
//
// zipFileName zip文件名 eg: test
func Zip(filePath string, zipFileName string) error {
	zipPath := fmt.Sprintf("%s\\%s.zip", filePath, zipFileName)
	zipfile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 遍历到的目录为当前目录则跳过
		if path == filePath {
			return err
		}

		// 遍历到的文件为压缩文件则跳过
		if info.Name() == (zipFileName + ".zip") {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		path = strings.ReplaceAll(path, "/", "\\")
		filePath = strings.ReplaceAll(filePath, "/", "\\")
		if !strings.HasSuffix(filePath, "\\") {
			filePath += "\\"
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(filePath)+"\\")

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

// Unzip zip解压缩
//
// zipFile zip文件路径 eg: /opt/test.zip
//
// toPath 	解压到目录 eg: /opt/test
func Unzip(zipFile string, toPath string) error {

	f, err := os.Open(zipFile)
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	reader, e := zip.NewReader(f, stat.Size())
	if nil != e {
		return e
	}

	// 确保解压目录存在
	if err := os.MkdirAll(toPath, 0666); err != nil {
		return err
	}

	var decodeName string
	for _, file := range reader.File {
		if file.Flags == 0 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i := bytes.NewReader([]byte(file.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			decodeName = string(content)
		} else {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			decodeName = file.Name
		}

		var fp = path.Join(toPath, decodeName)
		if file.FileInfo().IsDir() {
			if e := os.MkdirAll(fp, file.FileInfo().Mode()); nil != e {
				return e
			}
			continue
		}

		readcloser, e := file.Open()
		if nil != e {
			return e
		}

		b, e := ioutil.ReadAll(readcloser)
		if nil != e {
			return e
		}
		readcloser.Close()

		if e := ioutil.WriteFile(fp, b, file.FileInfo().Mode()); nil != e {
			return e
		}

	}
	return nil
}
