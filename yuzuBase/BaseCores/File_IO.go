package yuzuBase

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// writeIn 向文件写入数据
func writeIn(context string, path string) {
	//获取文件句柄
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(context)
	if err != nil {
		panic(err)
	}
}

// readOut 从文件读出数据
func readOut(path string) []string {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	tmpStr := make([]string, 0)
	for {
		//一行一行读取
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			break
		}
		tmpStr = append(tmpStr, line)
	}
	return tmpStr
}

// writeIni 根据路径写入数据库的配置信息,文件名为yuzu.ini
func writeIni(path string, db *DataBase) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	str := ""
	for i := 0; i < len(db.formName); i++ {
		if i == len(db.formName)-1 {
			str += db.formName[i] + "\n"
		} else {
			str += db.formName[i] + " "
		}
	}
	_, err = f.WriteString(str)
	if err != nil {
		panic(err)
	}
}

func readIni(path string) []string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	tmpStr := make([]string, 0)
	for {
		//一行一行读取
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			break
		}
		//去掉换行符
		line = line[:len(line)-1]
		tmpArr := strings.Fields(line)
		for _, v := range tmpArr {
			tmpStr = append(tmpStr, v)
		}

	}
	return tmpStr
}
