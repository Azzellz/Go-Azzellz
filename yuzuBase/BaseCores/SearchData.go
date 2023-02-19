package yuzuBase

import (
	"fmt"
	"strings"
)

func (df *dataForm) SearchById(yuzuId int, argu ...interface{}) (json string) {
	if yuzuId > df.yuzuLength-1 {
		fmt.Println("超过现有Id限制")
		return
	}

	flag := true
	if len(argu) != 0 {
		if val, ok := argu[0].(bool); ok {
			switch val {
			case true:
				flag = true
			case false:
				flag = false
			}
		}

	}

	json = df.toJSONString(df.dataBody[yuzuId], flag)
	return
}

func (df *dataForm) SearchByField(fieldName string, target interface{}, argu ...interface{}) (json string) {

	//空接口转化为字符串
	target = fmt.Sprintf("%v", target)

	flag := true
	if len(argu) != 0 {
		if val, ok := argu[0].(bool); ok {
			switch val {
			case true:
				flag = true
			case false:
				flag = false
			}
		}
	}

	fieldIndex := 0
	for i, v := range df.fieldName {
		if v == fieldName {
			fieldIndex = i
		}
	}

	for _, v := range df.dataBody {
		tmp := v[3:]
		valArr := strings.Fields(tmp)
		valArr[len(valArr)-1] = valArr[len(valArr)-1][:len(valArr[len(valArr)-1])-1]
		if valArr[fieldIndex] == target {
			json = df.toJSONString(v, flag)
			return
		}
	}
	json = "can't found"
	return
}

// SearchPosition 用于返回数据位于表体的索引数
func (df *dataForm) searchPosition(fieldName string, target interface{}) (bodyIndex int) {

	//如果没找到则返回-1
	bodyIndex = -1

	//空接口转化为字符串
	target = fmt.Sprintf("%v", target)

	fieldIndex := 0
	for i, v := range df.fieldName {
		if v == fieldName {
			fieldIndex = i
		}
	}

	for i, v := range df.dataBody {
		tmp := v[3:]
		valArr := strings.Fields(tmp)
		valArr[len(valArr)-1] = valArr[len(valArr)-1][:len(valArr[len(valArr)-1])-1]
		if valArr[fieldIndex] == target {
			bodyIndex = i
			return
		}
	}
	return
}
