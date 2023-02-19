package yuzuBase

import (
	"fmt"
	"os"
)

// DataBase 数据库模型
type DataBase struct {
	baseName    string               //记录数据库的名字
	formMap     map[string]*dataForm //记录所有数据表
	currentForm *dataForm            //记录当前数据库装载的数据表
	baseUrl     string               //记录当前数据库所在目录
	formName    []string             //记录当前数据库所包含所有的表名
}

// dataForm  数据表模型
type dataForm struct {
	//dataMap    map[string]interface{} //先留着
	yuzuLength int      //记录数据数并且为数据编号
	dataBody   []string //以字符串形式展示数据(不包含字段名和-):运行时操作按这个字段来操作,避免多次io操作
	fieldName  []string //记录字段
	dataHead   string   //数据表头
	formName   string   //数据表名
}

// MakeBase 构建数据库对象
func MakeBase(baseName string, path string) *DataBase {
	baseUrl := ""
	if path == "." {
		baseUrl = path + "/" + baseName
		err := os.Mkdir(baseUrl, 0777)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		baseUrl = path + baseName
		err := os.Mkdir(baseUrl, 0777)
		if err != nil {
			fmt.Println(err)
		}
	}

	return &DataBase{formMap: make(map[string]*dataForm), currentForm: nil, baseName: baseName, baseUrl: baseUrl, formName: make([]string, 0)}
}

// makeForm 构建数据表对象
func makeForm(formName string) *dataForm {
	return &dataForm{dataBody: make([]string, 0), yuzuLength: 0, fieldName: make([]string, 0), dataHead: "", formName: formName}
}

// backForm 返回字符串形式的数据表,用于结束时写入数据
func (df *dataForm) backForm() (strForm string) {
	strForm += df.dataHead
	for _, v := range df.dataBody {
		strForm += v
	}
	return
}

// OpenBase 根据传入的url打开数据库
func OpenBase(path string) (*DataBase, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 文件夹不存在
		return nil, err
	} else {
		// 文件夹存在
		//获取数据库的名字
		baseName := ""
		for i := len(path) - 1; i > 0; i-- {
			if path[i] == '/' {
				baseName = path[i+1:]
			}
		}
		formMap := make(map[string]*dataForm)

		formNames := readIni(path + "/yuzu.ini")

		for _, formName := range formNames {

			tmpForm := &dataForm{formName: formName}
			tmpForm.initForm(path)
			formMap[formName] = tmpForm
		}
		return &DataBase{baseName: baseName, baseUrl: path, currentForm: nil, formMap: formMap, formName: formNames}, nil
	}
}
