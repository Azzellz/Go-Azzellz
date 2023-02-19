package yuzuBase

import "strconv"

// AddData 为表添加数据:按照$ ... $格式
func (df *dataForm) AddData(data ...interface{}) {
	//格式化数据
	for k := 0; k < len(data); k++ {
		values := getValueInfo(data[k])
		formattedStr := "$" + strconv.Itoa(df.yuzuLength) + " "
		for i := 0; i < len(values); i++ {
			if i == len(values)-1 {
				formattedStr += values[i]
			} else {
				formattedStr += values[i] + " "
			}
		}
		formattedStr += "$\n"
		df.dataBody = append(df.dataBody, formattedStr)
		df.yuzuLength++
	}

}
