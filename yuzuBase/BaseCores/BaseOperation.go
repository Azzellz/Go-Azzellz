package yuzuBase

// MakeForm makeForm 创建根据字符串命名创建数据表
func (db *DataBase) MakeForm(formName string) {
	df := makeForm(formName)
	db.formName = append(db.formName, formName)
	db.formMap[formName] = df
}

// LoadForm 装载数据表,用于数据库对数据表的操作
func (db *DataBase) LoadForm(formName string) {
	db.currentForm = db.formMap[formName]
}

// RegisterForm 根据传入的结构体实例注册数据表
func (db *DataBase) RegisterForm(form interface{}) {
	db.currentForm.Register(form)
}

// EmptyForm 用于判断当前装载的表是否为空
func (db *DataBase) EmptyForm() bool {
	return db.currentForm.dataHead == ""
}

//将数据表对数据的操作封装进数据库:

// Add 为当前装载的表添加数据
func (db *DataBase) Add(data ...interface{}) {
	db.currentForm.AddData(data)
}

// Remove 为当前装载的表删除数据
func (db *DataBase) Remove(fieldName string, target interface{}) {
	db.currentForm.RemoveData(fieldName, target)
}

// SearchById 在当前装载的表中根据内嵌隐藏字段YuzuId查找数据,返回类Json格式的字符串
func (db *DataBase) SearchById(yuzuId int, argu ...interface{}) string {
	return db.currentForm.SearchById(yuzuId, argu)
}

// SearchByField 在当前装载的表中根据字段名查找数据,返回类Json格式的字符串
func (db *DataBase) SearchByField(fieldName string, target interface{}, argu ...interface{}) string {
	return db.currentForm.SearchByField(fieldName, target, argu)
}

// Update 为当前装载的表更新单个数据
func (db *DataBase) Update(fieldName string, target interface{}, replace string) {
	db.currentForm.UpdateData(fieldName, target, replace)
}

// UpdateAll 为当前装载的表更新整条数据
func (db *DataBase) UpdateAll(fieldName string, target interface{}, replace string) {
	db.currentForm.UpdateAllData(fieldName, target, replace)
}

// InitForm 使用当前装载的表读取path路径的数据,初始化表
func (db *DataBase) InitForm() {
	db.currentForm.initForm(db.baseUrl)
}

// CloseForm 关闭当前表,将内存中的表数据覆盖写入路径
func (db *DataBase) CloseForm() {
	db.currentForm.closeForm(db.baseUrl + "/" + db.currentForm.formName + ".txt")
	db.writeIni()
}

// ShowForm 以字符串的形式打印出当前数据库装载的表
func (db *DataBase) ShowForm() {
	db.currentForm.showForm()
}

// WriteIni 根据路径写入数据库的配置信息,文件名为yuzu.ini
func (db *DataBase) writeIni() {
	writeIni(db.baseUrl+"/yuzu.ini", db)
}

// readIni 读取配置文件
func (db *DataBase) readIni() {
	db.formName = readIni(db.baseUrl + "/yuzu.ini")
}
