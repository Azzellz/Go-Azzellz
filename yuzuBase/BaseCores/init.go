package yuzuBase

import (
	"strings"
)

// initForm 用于打开数据库后的初始化工作:返回一个初始化好的数据
func (df *dataForm) initForm(baseUrl string) {
	//默认表在数据库文件夹下创建
	path := baseUrl + "/" + df.formName + ".txt"
	formArr := readOut(path)
	//初始化yuzuLength
	df.yuzuLength = len(formArr) - 2
	//初始化字段名切片
	fieldStr := formArr[0][8:]
	fieldStr = fieldStr[:len(fieldStr)-2]
	fieldNames := strings.Fields(fieldStr)
	df.fieldName = fieldNames
	//初始化表头
	formHeadArr := formArr[:2]
	df.dataHead = formHeadArr[0] + formHeadArr[1]
	//初始化表体
	formBodyArr := formArr[2:]
	for _, v := range formBodyArr {
		df.dataBody = append(df.dataBody, v)
	}

}
