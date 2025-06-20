package WebSocket

//import (
//	"encoding/json"
//	"github.com/LJin-go/exam/global"
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//	"time"
//)
//
///**
//一对一聊天
//{
//    "cmd":"send",
//    "data":{
//        "to_user_id": 2,
//        "message":"hello 2"
//    }
//}
//
//群聊
//{
//    "cmd":"sendGroup",
//    "data":{
//        "message":"all"
//    }
//}
//*/
//
//type WSReq struct {
//	Cmd  string      `json:"cmd"`
//	Data interface{} `json:"data"`
//}
//
//type WSResp struct {
//	Code uint        `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data"`
//}
//
//type SendS struct {
//	ToUserId uint   `json:"to_user_id"`
//	Message  string `json:"message"`
//}
//
//type WSSendReq struct {
//	Cmd  string `json:"cmd"`
//	Data SendS  `json:"data"`
//}
//
//type GroupS struct {
//	Message string `json:"message"`
//}
//
//type WSGroupReq struct {
//	Cmd  string `json:"cmd"`
//	Data GroupS `json:"data"`
//}
//
//func Chat(c *gin.Context) {
//	//模拟token获取UserID
//	//get params test
//	//userIdStr, _ := c.GetQuery("userId")
//	//var userId uint
//	//userIdInt, _ := strconv.Atoi(userIdStr)
//	//userId = uint(userIdInt)
//
//	//get middle userId
//	userId := c.GetUint("userId")
//
//	// update websocket
//	var upgrader = websocket.Upgrader{
//		CheckOrigin: func(r *http.Request) bool {
//			return true // 任何来源都允许
//		},
//	} // use default options
//
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		log.Print("upgrade:", err)
//		c.JSON(http.StatusOK, gin.H{
//			"code": -1,
//			"msg":  "websocket fail",
//			"data": nil,
//		})
//		return
//	}
//
//	defer conn.Close()
//
//	go func() {
//		// 30s
//		for {
//			err = conn.WriteMessage(websocket.PingMessage, []byte{})
//			if err != nil {
//				return
//			}
//
//			time.Sleep(30 * time.Second)
//		}
//	}()
//
//	//online user data
//	global.OnlineUser[userId] = conn
//	log.Println(global.OnlineUser)
//	for {
//		_, message, err := conn.ReadMessage()
//		if err != nil {
//			log.Println("read:", err)
//			break
//		}
//
//		var req WSReq
//
//		err = json.Unmarshal(message, &req)
//		if err != nil {
//
//			WSRespErr(conn, 1000, "illegal json")
//			continue
//		}
//
//		switch req.Cmd {
//		case "send":
//			go Send(conn, message, userId)
//		case "sendGroup":
//			go SendGroup(conn, message, userId)
//		default:
//			WSRespErr(conn, 10001, "no function")
//		}
//
//	}
//}
//
//func SendGroup(conn *websocket.Conn, message []byte, userId uint) {
//	var req WSGroupReq
//	err := json.Unmarshal(message, &req)
//	if err != nil {
//		WSRespErr(conn, 20001, "Parsing json failed ")
//		return
//	}
//
//	if req.Data.Message == "" {
//		WSRespErr(conn, 20004, "not Message params")
//		return
//	}
//
//	if len(global.OnlineUser) == 0 {
//		WSRespErr(conn, 20002, "not Online user")
//		return
//	}
//
//	for _, onlineConn := range global.OnlineUser {
//
//		WSRespSuccess(onlineConn, req.Data.Message)
//
//	}
//}
//
//func Send(conn *websocket.Conn, message []byte, userId uint) {
//
//	var req WSSendReq
//	err := json.Unmarshal(message, &req)
//	if err != nil {
//		WSRespErr(conn, 20001, "Parsing json failed ")
//		return
//	}
//
//	if req.Data.ToUserId < 1 {
//		WSRespErr(conn, 20002, "not ToUserId params")
//		return
//	}
//
//	if req.Data.Message == "" {
//		WSRespErr(conn, 20004, "not Message params")
//		return
//	}
//
//	if global.OnlineUser[req.Data.ToUserId] == nil {
//		WSRespErr(conn, 20003, "not online")
//		return
//	}
//
//	WSRespSuccess(global.OnlineUser[req.Data.ToUserId], req.Data.Message)
//
//	WSRespSuccess(conn, "send ok")
//
//}
//
//func WSRespErr(conn *websocket.Conn, code uint, msg string) {
//	response := WSResp{
//		Code: code,
//		Msg:  msg,
//		Data: nil,
//	}
//
//	responseStr, _ := json.Marshal(response)
//
//	err := conn.WriteMessage(websocket.TextMessage, responseStr)
//	if err != nil {
//		log.Println("write:", err)
//	}
//}
//
//func WSRespSuccess(conn *websocket.Conn, data interface{}) {
//	response := WSResp{
//		Code: http.StatusOK,
//		Msg:  "ok",
//		Data: data,
//	}
//
//	responseStr, _ := json.Marshal(response)
//
//	err := conn.WriteMessage(websocket.TextMessage, responseStr)
//	if err != nil {
//		log.Println("write:", err)
//	}
//}
