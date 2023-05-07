package main

import (
	"Net_Programing/tcp_chat/protocal"
	"fmt"
	"net"
)

const (
	ServerTip = "[Server-Tip]"
)

// 枚举用户的选择

type Server struct {
	listener net.Listener
	clients  map[net.Addr]Client //记录当前在线的所有用户
	//msgQueue  chan protocal.Message //消息信道
	chatWorld ChatWorld
}

type Client struct {
	conn     net.Conn
	name     string
	chatting chan bool //标识用户是否退出聊天室
	outRoom  chan bool //标识用户是否在聊天室
}

type ChatWorld struct {
	chatRooms map[string]ChatRoom //映射多个聊天室
}

type ChatRoom struct {
	id      int
	addCh   chan Client
	delCh   chan Client
	msgCh   chan []byte
	clients map[net.Addr]Client //记录当前聊天室内在线的所有用户
}

func NewServer(listenAddr string) (s *Server) {
	fmt.Println(ServerTip, "Server started on", listenAddr)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	s = &Server{
		listener: ln,
		clients:  make(map[net.Addr]Client),
		//msgQueue:  make(chan protocal.Message),
		chatWorld: ChatWorld{chatRooms: make(map[string]ChatRoom)},
	}
	//选择动态生成聊天室的模式

	return
}

// BoardCast 向所有在线群体广播消息
func (s *Server) BoardCast(msg string) {
	for _, cl := range s.clients {
		_, err := cl.conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// BoardCastInRoom 广播传入roomName的聊天室内的所有连线用户
func (s *Server) BoardCastInRoom(roomName string, msg string) {
	for _, cl := range s.chatWorld.chatRooms[roomName].clients {
		_, err := cl.conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// BoardTarget 广播一位指定用户
func (s *Server) BoardTarget(client Client, msg string) {
	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
	}

}

func ReadBytes(conn net.Conn) []byte {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("%v connection: %v is closed\n", ServerTip, conn.RemoteAddr())
		conn.Close()
		return nil
	}
	return buf[:n]
}

// DealOrder 解析用户发给服务器的指令,返回一个标识是否为指令的布尔值
func (s *Server) DealOrder(roomName string, order string, client Client) bool {
	switch order {
	case protocal.ShowClientsSumInRoom: //打印出所有在线用户,只在发起指令的终端出打印
		clients := fmt.Sprintf("目前聊天室[%v]内的用户(%v):\n", roomName, len(s.clients))
		for _, cl := range s.chatWorld.chatRooms[roomName].clients {
			clients += fmt.Sprintf("「%v ~ %v」\n", cl.name, cl.conn.RemoteAddr().String())
		}
		s.BoardTarget(client, clients)
		return true
	case protocal.ShowHelpInfo: //打印可用的指令列表
		orderList := fmt.Sprintf("可用指令如下:\n!help:显示有效的指令列表\n!clients:显示当前聊天室所有的在线用户\n!name:显示自己当前的昵称\n!rooms:显示当前所有可用的聊天室\n!quit:退出当前聊天室\n!exit:直接退出客户端程序\n")
		s.BoardTarget(client, orderList)
		return true
	case protocal.ShowSelfName: //显示自己的聊天名称
		s.BoardTarget(client, fmt.Sprintf("你在当前聊天室的昵称为:%v", client.name))
		return true
	case protocal.ShowLiveRooms: //显示当前所有可用的聊天室的信息:包括人数,聊天室id
		roomsSum := len(s.chatWorld.chatRooms)
		rooms := fmt.Sprintf("目前可用的聊天室[%d]:\n", roomsSum)
		//这里不能保证有序遍历
		for roomName, room := range s.chatWorld.chatRooms {
			rooms += fmt.Sprintf("[id:%d]:%v\n", room.id, roomName)
		}
		s.BoardTarget(client, rooms)
		return true
	case protocal.QuitRoom: //退出当前聊天室,但是并没有与服务器断开连接,可以重新选择,同时向聊天室内的所有在线用户广播
		//提前写入退出标识
		client.chatting <- true

		//从当前聊天室map中删除
		delete(s.chatWorld.chatRooms[roomName].clients, client.conn.RemoteAddr())
		s.BoardCastInRoom(roomName, fmt.Sprintf("!!!「%v」退出聊天室[%v],ip:%v !!!", client.name, roomName, client.conn.RemoteAddr().String()))
		//打印服务器日志
		fmt.Printf("%v New client [%v] quit the chatroom [%v] \n", ServerTip, client.conn.RemoteAddr(), roomName)
		fmt.Printf("%v Now clients sum in the chatroom [%v] :  %v\n", ServerTip, roomName, len(s.chatWorld.chatRooms[roomName].clients))
		fmt.Printf("%v Now clients sum online: %v \n", ServerTip, len(s.clients))
		return true
	}

	return false
}

func (s *Server) BoardChatRoomsNum(client Client) {
	roomsSum := len(s.chatWorld.chatRooms)
	rooms := fmt.Sprintf("目前可用的聊天室[%d]:\n", roomsSum)
	//这里不能保证有序遍历
	for roomName, room := range s.chatWorld.chatRooms {
		rooms += fmt.Sprintf("[id:%d | %v人]:%v\n", room.id, len(room.clients), roomName)
	}
	s.BoardTarget(client, rooms)
}

// CheckRoomLive 检测目标Room是否存在
func (s *Server) CheckRoomLive(roomName string) bool {
	_, ok := s.chatWorld.chatRooms[roomName]
	return ok
}

// ChatRoomBackground 根据传入的聊天室名字开启聊天室后台
func (s *Server) ChatRoomBackground(roomName string) {
	go func() {
		for {
			select {
			case client := <-s.chatWorld.chatRooms[roomName].addCh: //有client发起tcp连接,即进入聊天室
				s.chatWorld.chatRooms[roomName].clients[client.conn.RemoteAddr()] = client
				//打印服务器日志
				fmt.Printf("%v New client [%v] login the chatroom [%v] \n", ServerTip, client.conn.RemoteAddr(), roomName)
				fmt.Printf("%v Now clients sum in the chatroom [%v] : %v", ServerTip, roomName, len(s.chatWorld.chatRooms[roomName].clients))
				fmt.Printf("%v Now clients sum online: %v \n", ServerTip, len(s.clients))
				//在当前聊天室广播上线用户
				s.BoardCastInRoom(roomName, fmt.Sprintf("!!!「%v」进入聊天室[%v],ip:%v !!!", client.name, roomName, client.conn.RemoteAddr().String()))
			case client := <-s.chatWorld.chatRooms[roomName].delCh: //有client退出tcp连接,即退出客户端
				delete(s.clients, client.conn.RemoteAddr())
				fmt.Printf("%v Client [%v] disconnected from the chatroom [%v]\n", ServerTip, client.conn.RemoteAddr(), roomName)
				fmt.Printf("%v Now clients sum in the chatroom [%v] : %v", ServerTip, roomName, len(s.chatWorld.chatRooms[roomName].clients))
				fmt.Printf("%v Now clients sum online: %v \n", ServerTip, len(s.clients))
				//在当前聊天室广播离线用户
				//先判断下用户是否还在聊天室,如果不在就不用再广播了
				if _, ok := s.chatWorld.chatRooms[roomName].clients[client.conn.RemoteAddr()]; ok {
					s.BoardCastInRoom(roomName, fmt.Sprintf("!!!「%v」退出聊天室[%v],ip:%v !!!", client.name, roomName, client.conn.RemoteAddr().String()))
					delete(s.chatWorld.chatRooms[roomName].clients, client.conn.RemoteAddr())
				}

			case msg := <-s.chatWorld.chatRooms[roomName].msgCh: //有client发消息
				//服务器给每个连接的client广播信息
				s.BoardCastInRoom(roomName, string(msg))
			}
		}
	}()
}

// CreateChatRoom 先创建聊天室之后再加入,返回标识结果的布尔值
func (s *Server) CreateChatRoom(roomName string, client Client) bool {
	//先查询是否已存在同名聊天室,若存在则广播给用户,客户端应该重试
	if s.CheckRoomLive(roomName) {
		s.BoardTarget(client, protocal.RepeatChatRoom)
		return false
	}
	//初始化聊天室
	s.chatWorld.chatRooms[roomName] = ChatRoom{
		id:      len(s.chatWorld.chatRooms),
		addCh:   make(chan Client),
		delCh:   make(chan Client),
		msgCh:   make(chan []byte),
		clients: make(map[net.Addr]Client),
	}
	//根据传入的聊天室的名称开启后台监听
	s.ChatRoomBackground(roomName) //注意!:聊天室后台需要被确保只能开启一次!!!!!!
	//加入聊天室,返回表示是否成功加入聊天室的布尔值
	return s.JoinChatRoom(roomName, client)
}

// JoinChatRoom 先判断是否有对应聊天室,没有则给用户广播
func (s *Server) JoinChatRoom(roomName string, client Client) bool {
	//检查是否存在目标聊天室
	if !s.CheckRoomLive(roomName) {
		return false //加入失败
	}

	//将用户实例送入addCh并且此时会被后台捕获到
	s.chatWorld.chatRooms[roomName].addCh <- client
	//开启一个协程处理连接,表示成功加入聊天室
	go func() {

		//获取目标client的输入并做出相应操作:如指令解析
		for {
			select {
			case <-client.chatting: //用户退出聊天室
				//给用户广播!quit使客户端的监听后台退出监听
				s.BoardTarget(client, protocal.QuitRoom)

				client.outRoom <- true //告知用户处于大厅中
				return
			default:
				tip := []byte(fmt.Sprintf("「%v ~ %v」:", client.name, client.conn.RemoteAddr().String()))
				buf := make([]byte, 4096)
				n, err := client.conn.Read(buf)
				if err != nil {
					//客户端断开连接
					s.chatWorld.chatRooms[roomName].delCh <- client
					client.conn.Close()
					return
				}
				//指令判断并处理
				if !s.DealOrder(roomName, string(buf[:n]), client) {
					//传入msg管道,广播给聊天室内的所有在线用户
					s.chatWorld.chatRooms[roomName].msgCh <- append(tip, buf[:n]...)
				}

			}
		}

	}()

	return true //加入成功
}

// All 聊天室综合化操作
func (s *Server) All() {
	//先监听连接,获取传来的聊天室名称和连接client的昵称
	//不断监听,发现一个连接就送入addCh,此时会被后台捕获
	//只要一进入程序,与服务器的连接就算建立成功,会被储存在服务器的clients map里
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//监听从客户端传来的初始化信息:如昵称,聊天室名称:
		go func() {

			//第一步先获取用户传来的昵称:
			username := string(ReadBytes(conn))
			//初始化用户实例
			client := Client{
				conn:     conn,
				name:     username,
				chatting: make(chan bool, 1),
				outRoom:  make(chan bool),
			}

			//连线成功则在server的map里注册client
			s.clients[conn.RemoteAddr()] = client

			//此时应当向用户发送当前聊天室的存在情况,并且让用户选择加入聊天室或是创建聊天室
			s.BoardChatRoomsNum(client)

			//接收用户的第二个输入,即用户选择加入聊天室或是创建聊天室,客户端应该发来两个对应枚举量
			var chatRoomName, userChoice string
			for {

				buf := make([]byte, 4096)
				n, err := client.conn.Read(buf)
				if err != nil {
					//客户端断开连接
					s.chatWorld.chatRooms[chatRoomName].delCh <- client
					client.conn.Close()
					return
				}
				userChoice = string(buf[:n])

				//fmt.Println(username, chatRoomName, userChoice)
				switch userChoice {
				case protocal.JoinChatRoom:
					chatRoomName = string(ReadBytes(conn))
					for !s.JoinChatRoom(chatRoomName, client) {
						//等待用户重新操作,此时客户端应该让客户重新输入要创建的聊天室
						chatRoomName = string(ReadBytes(conn))
					}
				case protocal.CreateChatRoom:
					chatRoomName = string(ReadBytes(conn))
					for !s.CreateChatRoom(chatRoomName, client) {
						//等待用户重新操作,此时客户端应该让客户重新输入要加入的聊天室
						chatRoomName = string(ReadBytes(conn))
					}
				case protocal.ShowLiveRooms:
					s.BoardChatRoomsNum(client)
					continue
				}

				<-client.outRoom //阻塞,能通过说明此时用户处于大厅,开启下一轮循环
				if string(ReadBytes(client.conn)) == protocal.QuitRoomSuccess {
					s.BoardChatRoomsNum(client)
				}
			}

		}()

	}
}

//func (s *Server) StartChatRoom() {
//	//开启后台监听
//	go func() {
//		for {
//			select {
//			case client := <-s.chatRoom.addCh: //有client发起tcp连接
//				s.clients[client.conn.RemoteAddr()] = client
//				fmt.Println(ServerTip, "New client connected:", client.conn.RemoteAddr())
//				fmt.Println(ServerTip, "Now clients sum:", len(s.clients))
//				//广播上线用户
//				s.BoardCast(fmt.Sprintf("!!!「%v」进入聊天室,ip:%v !!!", client.name, client.conn.RemoteAddr().String()))
//			case client := <-s.chatRoom.delCh: //有client退出tcp连接
//				delete(s.clients, client.conn.RemoteAddr())
//				fmt.Println(ServerTip, "Client disconnected:", client.conn.RemoteAddr())
//				fmt.Println(ServerTip, "Now clients sum:", len(s.clients))
//				//广播离线用户
//				s.BoardCast(fmt.Sprintf("!!!「%v」退出聊天室,ip:%v !!!", client.name, client.conn.RemoteAddr().String()))
//			case msg := <-s.chatRoom.msgCh: //有client发消息
//				//服务器给每个连接的client广播信息
//				s.BoardCast(string(msg))
//			}
//		}
//	}()
//
//	//不断监听,发现一个连接就送入addCh,此时会被后台捕获
//	for {
//		conn, err := s.listener.Accept()
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		//监听从客户端传来的初始化信息:如昵称..
//		buf := make([]byte, 4096)
//		n, err := conn.Read(buf)
//		if err != nil {
//			return
//		}
//
//		//将初始化Client实例写入管道
//		client := Client{
//			conn: conn,
//			name: string(buf[:n]),
//		}
//		s.chatRoom.addCh <- client
//
//		//开启一个协程处理连接
//		go func() {
//			//处理完请求(目标客户端退出)后让后台捕获到并且断开连接
//			defer func() {
//				s.chatRoom.delCh <- client
//				client.conn.Close()
//			}()
//
//			//获取目标client的输入
//			for {
//				msg := []byte(fmt.Sprintf("「%v ~ %v」:", client.name, client.conn.RemoteAddr().String()))
//				buf := make([]byte, 4096)
//				n, err := conn.Read(buf)
//				if err != nil {
//					return
//				}
//				//指令判断并处理
//				if !s.DealOrder(string(buf[:n]), conn.RemoteAddr()) {
//					s.chatRoom.msgCh <- append(msg, buf[:n]...)
//				}
//
//			}
//		}()
//	}
//
//}

// Over 关闭服务器
func (s *Server) Over() {
	s.listener.Close()
}

func main() {
	s := NewServer("localhost:8080")
	defer s.Over()
	s.All()
}
