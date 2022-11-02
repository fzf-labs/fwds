package util

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var File = newFile()

type file struct {
}

func newFile() *file {
	return &file{}
}

// IsFileExists 文件是否存在
func (f *file) IsFileExists(path string) (file os.FileInfo, IsExist bool) {
	file, err := os.Stat(path)
	return file, err == nil || os.IsExist(err)
}

// CreateDir 批量创建文件夹
func (f *file) CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		_, exist := f.IsFileExists(v)
		if err != nil {
			return err
		}
		if !exist {
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				return err
			}
			if err := os.Chmod(v, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return err
}

// ReadDirAll 读取目录
// example ReadDirAll("/Users/why/Desktop/go/test", 0)
func (f *file) ReadDirAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			f.ReadDirAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}

// WriteWithIo 使用io.WriteString()函数进行数据的写入，不存在则创建
func (f *file) WriteWithIo(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	if content != "" {
		_, err := io.WriteString(file, content)
		if err != nil {
			return err
		}
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.", content)
	}

	return nil
}

// ReadLimit 读取指定字节
func (f *file) ReadLimit(str string, len int64) string {
	reader := strings.NewReader(str)
	limitReader := &io.LimitedReader{R: reader, N: len}

	var res string
	for limitReader.N > 0 {
		tmp := make([]byte, 1)
		_, err := limitReader.Read(tmp)
		if err != nil {
			return ""
		}
		res += string(tmp)
	}
	return res
}

// ReadFile 读取整个文件
func (f *file) ReadFile(dir string) (string, error) {
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadFileLine  按行读取文件
func (f *file) ReadFileLine(dir string) (map[int]string, error) {
	file, err := os.OpenFile(dir, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	/* stat, err := file.Stat()
	util_err.Must(err)
	size := stat.Size */

	buf := bufio.NewReader(file)
	res := make(map[int]string)
	i := 0
	for {
		line, _, err := buf.ReadLine()
		context := string(line)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		res[i] = context
		i++
	}
	return res, nil
}

// ReadJsonFile 读取json文件
func (f *file) ReadJsonFile(dir string) (string, error) {
	jsonFile, err := os.Open(dir)
	if err != nil {
		return "", err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return string(byteValue), nil
}

// GetFileInfo 获得文件Info
func (f *file) GetFileInfo(file *os.File) os.FileInfo {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("file stat error:", err)
	}
	return fileInfo
}

// GetFileMode 获得文件权限Mode
func (f *file) GetFileMode(file *os.File) os.FileMode {
	fileInfo := f.GetFileInfo(file)
	return fileInfo.Mode()
}

// Chown 更改文件所有者
func (f *file) Chown(file *os.File, uid, gid int) error {
	if uid == 0 {
		uid = os.Getuid()
	}
	if gid == 0 {
		gid = os.Getgid()
	}

	return file.Chown(uid, gid)
}

// Chmod 更改文件权限
func (f *file) Chmod(file *os.File, mode int) error {
	return file.Chmod(os.FileMode(mode))
}

// Open 打开文件
func (f *file) Open(dir string) (*os.File, error) {
	return os.Open(dir)
}

// Create 创建文件
func (f *file) Create(dir string) (*os.File, error) {
	return os.Create(dir)
}

// CleanFile 清楚文件内容
func (f *file) CleanFile(filePath string) error {
	fl, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer fl.Close()

	_, err = fl.WriteString("")
	return err
}
