package yuzuBase

import (
	"fmt"
	"strconv"
	"strings"
)

// UpdateData 更新单个数据
func (df *dataForm) UpdateData(fieldName string, target interface{}, replace string) {
	bodyIndex := df.searchPosition(fieldName, target)
	targetStr := fmt.Sprintf("%v", target)
	df.dataBody[bodyIndex] = strings.Replace(df.dataBody[bodyIndex], targetStr, replace, -1)
}

// UpdateAllData 更新整条数据
func (df *dataForm) UpdateAllData(fieldName string, target interface{}, replace string) {
	bodyIndex := df.searchPosition(fieldName, target)
	replace = "$" + strconv.Itoa(bodyIndex) + " " + replace + "$\n"
	df.dataBody[bodyIndex] = replace
}
