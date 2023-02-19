package yuzuBase

func (df *dataForm) RemoveData(fieldName string, target interface{}) {
	bodyIndex := df.searchPosition(fieldName, target)
	df.dataBody = append(df.dataBody[:bodyIndex], df.dataBody[bodyIndex+1:]...)
}
