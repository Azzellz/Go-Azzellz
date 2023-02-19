package yuzuBase

// Register 注册数据表(根据form样式)该方法只需要在注册表的时候调用一次
func (df *dataForm) Register(form interface{}) {

	//获取字段和值的两种切片
	fieldsArr := getFieldInfo(form)
	//fmt.Println(fieldsArr)
	//根据字段切片拼接字符串
	formHead := "?YuzuId "
	for i := 0; i < len(fieldsArr); i++ {
		if i == len(fieldsArr)-1 {
			formHead += fieldsArr[i]
		} else {
			formHead += fieldsArr[i] + " "
		}

	}
	formHead += "?\n-\n"
	df.dataHead = formHead
	df.fieldName = fieldsArr

}
