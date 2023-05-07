package main

import (
	"Net_Programing/tcp_chat/protocal"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	targetAddr string
	chatRoom   ChatRoom //维护聊天室的功能
	conn       net.Conn
	name       string
}

type ChatRoom struct {
	msg string
}

func ErrorTip(tip string) {
	for {
		fmt.Println(tip)
		fmt.Println("输入任意内容退出(除回车之外)")
		var something string
		fmt.Scanln(&something)
		if something != "" {
			os.Exit(1)
		}
	}
}

// NewClient 返回一个与目标host建立tcp连接的Client对象
func NewClient(targetAddr string, connectMode string) (c *Client) {
	conn, err := net.Dial(connectMode, targetAddr)
	if err != nil {
		fmt.Println(err)
		ErrorTip("连接服务器错误,请重试")
		return nil
	}
	c = &Client{
		targetAddr: targetAddr,
		conn:       conn,
		chatRoom:   ChatRoom{msg: ""},
		name:       "unknown",
	}
	return
}

func (c *Client) Over() {
	c.conn.Close()
}

func (c *Client) JoinChat() {

	//开一个后台监听接收服务器广播其他人发的信息
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := c.conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				ErrorTip("服务器已断开连接或服务器已关闭...")
			}
			if string(buf[:n]) == protocal.QuitRoom {
				//结束监听
				//给服务器发送成功退出
				c.WriteToServer(protocal.QuitRoomSuccess)
				return
			}
			//做个小优化,避免客户端看到重复的信息
			if c.chatRoom.msg == strings.TrimSpace(string(buf[:n])) {
				c.chatRoom.msg = ""
				continue
			}
			fmt.Println(strings.TrimSpace(string(buf[:n])))
			//fmt.Println(123)
		}
	}()

	//获取用户输入,并发送给服务器
	for {
		fmt.Scanln(&c.chatRoom.msg)
		//输入!exit直接退出
		if c.chatRoom.msg == protocal.Quit {
			os.Exit(1)
		}

		_, err := c.conn.Write([]byte(c.chatRoom.msg))
		if err != nil {
			fmt.Println(err)
			ErrorTip("与服务器断开连接...")
			return
		}

		//如果是!quit,要先通知服务器,再在客户端这里响应
		if c.chatRoom.msg == protocal.QuitRoom {
			return
		}

	}
}

func (c *Client) WriteToServer(msg string) {
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("error!")
		ErrorTip("写入信息失败,或许是与服务器断开连接?")
	}
}

func (c *Client) Guide() {
	fmt.Println("Welcome to YuzuChat !\n请输入你的昵称:")
	var username string
	fmt.Scanln(&username)
	//给服务器发送自己的昵称
	c.WriteToServer(username)

	fmt.Printf("Hello!  %v\n ", username)
	//在这里接收服务器提供的可用聊天室列表

	for {
		//这里是要获取在线的聊天室列表
		buf := make([]byte, 4096)
		n, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			ErrorTip("服务器已断开连接或服务器已关闭...")
		}
		fmt.Println(strings.TrimSpace(string(buf[:n])))
		//给出操作选项
		fmt.Println("请选择:\n1.创建聊天室\n2.加入聊天室\n3.查看可用聊天室\n0.退出聊天室")
		var choice int
		var roomName string

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			c.WriteToServer(protocal.CreateChatRoom)
			fmt.Println("请输入您要创建的聊天室的名称:")
			fmt.Scanln(&roomName)
			c.WriteToServer(roomName)
		case 2:
			c.WriteToServer(protocal.JoinChatRoom)
			fmt.Println("请输入您要加入的聊天室的名称:")
			fmt.Scanln(&roomName)
			c.WriteToServer(roomName)
		case 3:
			c.WriteToServer(protocal.ShowLiveRooms)
			continue
		case 0:
			os.Exit(1)
		default:
			fmt.Println("BYD,请重新输入")
		}

		c.JoinChat()
	}

}

func main() {
	cl := NewClient("a2528447065.e2.luyouxia.net:20904", "tcp")
	defer cl.Over()
	cl.Guide()
}
