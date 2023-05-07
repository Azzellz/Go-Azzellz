package protocal

// 服务器提示信息
const (
	RepeatChatRoom = "已经有重名聊天室,请重试"
)

// 客户端选择标识信息
const (
	JoinChatRoom    = "JoinChatRoom"
	CreateChatRoom  = "CreateChatRoom"
	QuitRoomSuccess = "QuitRoomSuccess"
)

// 客户端指令
const (
	ShowClientsSumInRoom = "!clients"
	ShowHelpInfo         = "!help"
	ShowSelfName         = "!name"
	ShowLiveRooms        = "!rooms"

	QuitRoom = "!quit"
	Quit     = "!exit" //退出客户端
)

//type Message struct {
//	MsgType int
//	Data    string
//	Origin  net.Addr //源地址
//}
//
//func NewMessage(msgType int, data string, origin net.Addr) *Message {
//	if msgType >= Max || msgType <= Min {
//		return nil
//	}
//	return &Message{MsgType: msgType, Data: data, Origin: origin}
//}
