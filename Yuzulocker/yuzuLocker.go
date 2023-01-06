package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {

	for {

		fmt.Println("Starting..")
		fmt.Println("使用说明:只需要输入路径,yuzu会自动判断是需要加密还是需要解密,加密文件的格式是op:")
		fmt.Println("请输入路径:")
		var path string
		fmt.Scan(&path)
		_, path = filepath.Split(path)
		var nowDir, firstName string
		for i := range path {
			if path[i] == '.' {
				nowDir = path[i+1:]
				if nowDir[len(nowDir)-1] == '"' { //说明是window直接复制过去的地址
					nowDir = nowDir[:len(nowDir)-1]
				}
				firstName = path[:i]
				break
			} //拿到一开始的后缀
		}
		fmt.Println(firstName, nowDir)
		//判断是加密还是解密:
		if nowDir == "op" {
			//进行解密
			path = "./" + firstName + "." + nowDir

			sF, err := os.Open(path)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			//抢在创建dF之前先获取后缀名
			f, _ := os.Open(path)

			reader := bufio.NewReader(f)
			//循环读取文件内容
			var suffix string
			for {
				str, err := reader.ReadString('\n') //读到换行符就结束
				//fmt.Printf(str)
				if err == io.EOF {
					suffix = str

					break
				}
			}

			err = f.Close()
			if err != nil {
				fmt.Printf("%v", err)
			}

			//找到后缀suffix
			fmt.Println(suffix)

			dfPath := "./" + firstName + "." + suffix
			dF, err := os.Create(dfPath)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			buf := make([]byte, 4*1024) //4k大小的临时缓冲区
			for {
				n, err0 := sF.Read(buf) //读取
				if err0 != nil {
					if err0 == io.EOF { //读取完毕后,io会报错,因此这个错误可以理解为读取完毕的象征
						break
					}
					fmt.Println("完毕")
				}
				buf[0] -= 1
				fmt.Println(n, len(suffix), string(buf[n-len(suffix):]), suffix)
				if string(buf[n-len(suffix):])[:len(suffix)] == suffix {
					fmt.Println("yes")
					dF.Write(buf[:n-len(suffix)-1])
					break
				}
				//
				dF.Write(buf[:n])

			}
			fmt.Println("解密完成!")
			dF.Close()
			sF.Close()

			fmt.Println("完毕,输入1退出....")
			fmt.Println("输入其他继续操作....")
			var a string
			fmt.Scan(&a)
			if a == "1" {
				break
			}

		} else { //加密
			sF, err := os.Open(path)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			dF, err := os.Create("./" + firstName + ".op")

			buf := make([]byte, 4*1024) //4k大小的临时缓冲区

			for {
				n, err0 := sF.Read(buf) //读取
				if err0 != nil {
					if err0 == io.EOF { //读取完毕后,io会报错,因此这个错误可以理解为读取完毕的象征
						dF.WriteString("\n")
						break
					}
				}

				//将后缀名按照某一算法进行加密

				buf[0] += 1
				dF.Write(buf[:n])
				fmt.Println(n)
			}
			_, err = dF.WriteString(nowDir)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("加密完成!")
			dF.Close()
			sF.Close()
			fmt.Println("完毕,输入1退出....")
			fmt.Println("输入其他继续操作....")
			var a string
			fmt.Scan(&a)
			if a == "1" {
				break
			}

		}

	}

}
