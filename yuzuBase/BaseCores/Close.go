package yuzuBase

// closeForm 用于结束数据表操作后的写入数据操作
func (df *dataForm) closeForm(path string) {
	writeIn(df.backForm(), path)
}
