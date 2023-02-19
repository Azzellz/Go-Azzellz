// 用于测试数据库功能
package main

import (
	yuzuBase "awesomeProject2/yuzuBase/BaseCores"
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
	sex  string
	fat  bool
}
type user2 struct {
	name string
}

func CreateStruct(s interface{}) interface{} {
	// 通过反射获取传入的结构体类型
	t := reflect.TypeOf(s)
	// 创建一个新的结构体类型
	v := reflect.New(t).Interface()
	// 将传入的结构体赋值给新的结构体
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(s))
	return v
}

func main() {
	//db := yuzuBase.MakeBase("shit", "./")
	//
	//db.MakeForm("shit")
	//
	//db.LoadForm("shit")
	////db.InitForm() //相当于读档操作
	//db.RegisterForm(user{})
	//tmpU := user{name: "shabiii",
	//	age: 10,
	//	sex: "male",
	//	fat: false}
	//db.Add(tmpU)
	//db.ShowForm()
	//db.CloseForm()
	//fmt.Println()

	db, err := yuzuBase.OpenBase("./shit")
	if err != nil {
		fmt.Println(err)
	}
	db.LoadForm("shit")
	tmpU := user{name: "shabiii",
		age: 10,
		sex: "male",
		fat: false}
	db.Add(tmpU)
	db.ShowForm()
	db.CloseForm()
	//db.LoadForm("homo")
	//db.RegisterForm(user2{})
	//db.Add(user2{name: "shit"})
	//db.ShowForm()
	//fmt.Println()
	//
	//db.CloseForm()

	//u1 := user{"312", 14, "123", true}
	//u2 := user{"3122", 144, "1123", false}
	//db := cores.makeForm()
	//db.Register(user{})
	//db.AddData(u1)
	//db.AddData(u2)
	//db.showForm()
	//
	//db.showForm()
	//cores.closeForm(db)

}
