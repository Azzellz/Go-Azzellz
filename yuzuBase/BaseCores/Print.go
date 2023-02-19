package yuzuBase

import "fmt"

// showForm 打印数据表
func (df *dataForm) showForm() {
	fmt.Print(df.dataHead)
	for _, v := range df.dataBody {
		fmt.Print(v)
	}
}
