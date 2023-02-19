package yuzuBase

import (
	"fmt"
	"reflect"
	"strings"
)

// GetInfo 将传入的结构体转换为带有字段和值信息的字符串
func getInfo(s interface{}) (valuesArr []string, fieldsArr []string) {
	valuesArr = getValueInfo(s)
	fieldsArr = getFieldInfo(s)
	return
}
func getValueInfo(s interface{}) (valuesArr []string) {
	// 通过反射获取传入的结构体类型
	t := reflect.TypeOf(s)
	// 创建一个新的结构体类型
	v := reflect.New(t).Interface()
	// 将传入的结构体赋值给新的结构体
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(s))
	//将包含信息的空接口转化为字符串
	tmpStr := fmt.Sprintf("%v", v)[3:]
	tmpStr = tmpStr[:len(tmpStr)-2]
	//分割字符串
	valuesArr = strings.Fields(tmpStr)
	return
}

func getFieldInfo(s interface{}) (fieldsArr []string) {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tmpArr := f.Name
		fieldsArr = append(fieldsArr, tmpArr)
	}
	return
}

// toJSONString 本方法用于生成json格式的数据,根据db实例中的字段名切片
func (df *dataForm) toJSONString(str string, mode bool) (json string) {
	if mode {
		str = str[1:]

		valArr := strings.Fields(str)
		valArr[len(valArr)-1] = valArr[len(valArr)-1][:len(valArr[len(valArr)-1])-1]
		json += "{YuzuId:" + valArr[0] + " "
		valArr = valArr[1:]
		for i, v := range df.fieldName {
			if i == len(df.fieldName)-1 {
				json += v + ":" + valArr[i] + "}"
			} else {
				json += v + ":" + valArr[i] + " "
			}
		}
	} else {
		str = str[3:]
		valArr := strings.Fields(str)
		valArr[len(valArr)-1] = valArr[len(valArr)-1][:len(valArr[len(valArr)-1])-1]
		fmt.Println(valArr)
		json += "{"
		for i, v := range df.fieldName {
			if i == len(df.fieldName)-1 {
				json += v + ":" + valArr[i] + "}"
			} else {
				json += v + ":" + valArr[i] + " "
			}
		}
	}

	return
}
