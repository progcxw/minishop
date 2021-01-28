package controllers

import (
	"encoding/json"
	"fmt"
	"minishop/services"
	"minishop/utils"
	"net/http"
	"sort"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/websocket"

	"minishop/models"
)

type ChatController struct {
	beego.Controller
}

//连接的客户端,把每个客户端对应其用户id
var chatMapper = make(map[int]*websocket.Conn)

// 配置升级程序(升级为websocket)
var upgrader = websocket.Upgrader{}

//OnOpen 客户端创建websocket链接
func (this *ChatController) OnOpen() {
	openId := this.GetString("open_id")
	o := orm.NewOrm()
	var user models.MinishopUser
	o.QueryTable("minishop_user").Filter("weixin_openid", openId).One(&user)

	// 解决跨域问题(微信小程序)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	//升级将HTTP服务器连接升级到WebSocket协议。
	//responseHeader包含在对客户端升级的响应中
	//请求。使用responseHeader指定Cookie（设置Cookie）和
	//应用程序协商的子目录（Sec WebSocket协议）。
	//如果升级失败，则升级将向客户端答复一个HTTP错误
	conn, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		fmt.Println("open websocket error: ", err)
		return
	}

	fmt.Println("open websocket success, userID = ", user.Id)
	chatMapper[user.Id] = conn

	unreadCount := getUnreadCount(user.Id)
	wsMessage := models.WsMessage{
		MessageType: utils.UnreadNum,
		MessageBody: fmt.Sprint(unreadCount),
		SendTime:    utils.GetTimestamp(),
	}

	rtnMsg := make(map[string]interface{})
	rtnMsg["errno"] = 0
	rtnMsg["errmsg"] = ""
	rtnMsg["data"] = wsMessage

	err = conn.WriteJSON(rtnMsg)
	if err != nil {
		fmt.Println("openws send message failed: ", err)
		return
	}

	var msg models.WsMessage
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("reading message from ", user.Nickname, " err: ", err)
			return
		}

		json.Unmarshal(message, &msg)

		fmt.Println("receive message from ", user.Nickname, " message: ", msg)
		sendMessage(msg.SenderId, msg.ReceiverId, msg.MessageType, msg.ChatId, msg.MessageBody)
	}
}

//OnClose 客户端关闭websocket链接
func (this *ChatController) OnClose() {
	openId := this.GetString("open_id")
	o := orm.NewOrm()
	var user models.MinishopUser
	o.QueryTable("minishop_user").Filter("weixin_openid", openId).One(&user)

	if conn, ok := chatMapper[user.Id]; ok {
		conn.Close()
	}

	delete(chatMapper, user.Id)
}

//CreateChat 创建聊天
func (this *ChatController) CreateChat() {
	senderId := getLoginUserId()

	receiverId, err := this.GetInt("receiver_id")
	if err != nil {
		fmt.Println("get receiver id error: ", err)
		return
	}
	goodsId, err := this.GetInt("goods_id")
	if err != nil {
		fmt.Println("get goods id error: ", err)
		return
	}

	o := orm.NewOrm()
	newChat := new(models.Chat)
	var chatId int64

	if senderId > receiverId {
		newChat.U1 = senderId
		newChat.U2 = receiverId
		newChat.GoodsId = goodsId
		newChat.ShowToU1 = true
		newChat.ShowToU2 = false
	} else {
		newChat.U1 = receiverId
		newChat.U2 = senderId
		newChat.GoodsId = goodsId
		newChat.ShowToU1 = false
		newChat.ShowToU2 = true
	}

	exist := o.QueryTable(newChat).Filter("U1", newChat.U1).Filter("U2", newChat.U2).Filter("goods_id", newChat.GoodsId).Exist()
	if exist == false {
		chatId, err = o.Insert(newChat)
		if err != nil {
			fmt.Println("chat insert error: ", err)
			return
		}

		var u1tou2 bool
		if senderId > receiverId {
			u1tou2 = true
		} else {
			u1tou2 = false
		}

		history := models.History{
			ChatId:      chatId,
			U1ToU2:      u1tou2,
			MessageType: utils.EstablishChat,
		}

		o.Insert(&history)
	}

	utils.ReturnHTTPSuccess(&this.Controller, chatId)
	this.ServeJSON()
}

//GetChatIndex 获取用户参与的chat中最后的消息,写入chatIndexEle中
func (this *ChatController) GetChatIndex() {
	userId := getLoginUserId()
	var chats []models.Chat

	//获取本人参与的聊天的chat
	o := orm.NewOrm()
	cond := orm.NewCondition()
	condition := cond.AndCond(cond.And("u1", userId).And("show_to_u1", true)).OrCond(cond.And("u2", userId).And("show_to_u2", true))
	num, err := o.QueryTable("chat").SetCond(condition).All(&chats)
	if err != nil {
		fmt.Println("getChatIndex get chat error: ", err)
		return
	}

	chatEles := make(models.ChatEleList, num, num)
	for i, singleChat := range chats {
		chatEles[i].LastChat.ChatId = singleChat.Id
		chatEles[i].LastChat.MessageBody = singleChat.LastMessage
		chatEles[i].LastChat.MessageType = utils.UserMessage
		chatEles[i].LastChat.U1ToU2 = singleChat.U1ToU2
		chatEles[i].LastChat.SendTime = singleChat.LastTime
		chatEles[i].GoodsId = singleChat.GoodsId
		chatEles[i].UnreadNum = singleChat.UnreadNum

		//获取chatEle的otherside
		var otherSideId int
		var otherSideUser models.MinishopUser
		if singleChat.U1 == userId {
			otherSideId = singleChat.U2
		} else {
			otherSideId = singleChat.U1
		}
		o.QueryTable("minishop_user").Filter("id", otherSideId).One(&otherSideUser)

		chatEles[i].OtherSide.NickName = otherSideUser.Nickname
		chatEles[i].OtherSide.AvatarUrl = otherSideUser.Avatar
		chatEles[i].OtherSide.UserId = otherSideId

		//获取本次聊天相关的商品信息
		var goodsinfo models.MinishopPostGoods
		o.QueryTable("minishop_post_goods").Filter("id", singleChat.GoodsId).One(&goodsinfo)

		chatEles[i].Goods.Id = singleChat.GoodsId
		chatEles[i].Goods.Name = goodsinfo.Name
		chatEles[i].Goods.Price = goodsinfo.Price
		primaryPicUrl := strings.Split(goodsinfo.ListPicUrl, " ")
		chatEles[i].Goods.PrimaryPicUrl = primaryPicUrl[0]
	}

	sort.Sort(chatEles)

	utils.ReturnHTTPSuccess(&this.Controller, models.ChatIndex{
		Chats:      chatEles,
		OffsetTime: utils.GetTimestamp(),
	})

	this.ServeJSON()
}

//GetChatForm 获取聊天框信息
func (this *ChatController) GetChatForm() {
	o := orm.NewOrm()
	userId := getLoginUserId()
	chatId, err := this.GetInt("chat_id")
	if err != nil {
		fmt.Println("getChatForm get chatId error:", err)
		return
	}

	//填写chatForm信息
	var chatForm models.ChatForm
	var chat models.Chat
	var otherSideUser models.MinishopUser
	var goodsInfo models.MinishopPostGoods
	o.QueryTable("chat").Filter("id", chatId).One(&chat)
	if userId == chat.U1 {
		chatForm.IsU1 = true
		chatForm.OtherSide.UserId = chat.U2
	} else {
		chatForm.IsU1 = false
		chatForm.OtherSide.UserId = chat.U1
	}

	o.QueryTable("minishop_user").Filter("id", chatForm.OtherSide.UserId).One(&otherSideUser)
	chatForm.OtherSide.NickName = otherSideUser.Nickname
	chatForm.OtherSide.AvatarUrl = otherSideUser.Avatar

	o.QueryTable("minishop_post_goods").Filter("id", chat.GoodsId).One(&goodsInfo)
	chatForm.Goods.Id = goodsInfo.Id
	chatForm.Goods.Name = goodsInfo.Name
	chatForm.Goods.Price = goodsInfo.Price
	primaryPicUrl := strings.Split(goodsInfo.ListPicUrl, " ")
	chatForm.Goods.PrimaryPicUrl = primaryPicUrl[0]

	//将已读信息填入chaForm
	o.QueryTable("history").Filter("chat_id", chatId).All(&chatForm.HistoryList)

	//将未读消息填入chatForm，并将自己未读的消息设为已读
	conn := services.RedisPool().Get()
	data, err := conn.Do("GET", chatId)
	if err != nil {
		fmt.Println("getChatForm getting unread message error: ", err)
		return
	}

	if data != nil {
		var unreadMsg []models.WsMessage
		err = json.Unmarshal(data.([]uint8), &unreadMsg)
		if err != nil {
			fmt.Println("chatForm json unmarshal error: ", err)
			return
		}

		unreadHisList := models.WsListToHisList(unreadMsg)
		chatForm.HistoryList = append(chatForm.HistoryList, unreadHisList...)

		if unreadMsg[0].ReceiverId == userId {
			conn.Do("DEL", chatId)
			o.QueryTable("chat").Filter("id", chatId).Update(orm.Params{"unread_num": 0})
		}
	}

	utils.ReturnHTTPSuccess(&this.Controller, chatForm)
	this.ServeJSON()
}

//getUnreadCount 获取未读消息数量
func getUnreadCount(userId int) int {
	var chatList []models.Chat
	var unreadCount int = 0

	o := orm.NewOrm()
	cond := orm.NewCondition()
	condition := cond.AndCond(cond.And("u1", userId).And("show_to_u1", true)).OrCond(cond.And("u2", userId).And("show_to_u2", true))
	o.QueryTable("chat").SetCond(condition).All(&chatList)

	for _, chat := range chatList {
		if userId == chat.U1 {
			if chat.U1ToU2 == false {
				unreadCount += chat.UnreadNum
			}
		} else {
			if chat.U1ToU2 == true {
				unreadCount += chat.UnreadNum
			}
		}
	}

	return unreadCount
}

//sendMessage 发送消息到对应的websocket
func sendMessage(senderId, receiverId, messageType int, chatId int64, message string) {
	o := orm.NewOrm()
	sendTime := utils.GetTimestamp()

	m := models.WsMessage{
		ChatId:      chatId,
		ReceiverId:  receiverId,
		SenderId:    senderId,
		MessageBody: message,
		MessageType: messageType,
		SendTime:    sendTime,
	}

	var u1tou2 bool
	if senderId > receiverId {
		u1tou2 = true
	} else {
		u1tou2 = false
	}

	if wsConn, ok := chatMapper[receiverId]; !ok {
		//将未读消息加入redis
		fmt.Println("sendmessage didn't find a live conn, receiver id= ", receiverId)
		conn := services.RedisPool().Get()
		unreadMsg, err := json.Marshal(m)
		if err != nil {
			fmt.Println("sendmessage json marshal error: ", err)
			return
		}

		_, err = conn.Do("APPEND", chatId, unreadMsg)
		if err != nil {
			fmt.Println("redis unread message set error: ", err)
			return
		}
	} else {
		fmt.Println("发送成功")
		err := wsConn.WriteJSON(m)
		if err != nil {
			fmt.Println("sendmessage failed, error: ", err)
			return
		}

		//将已读消息写入mysql的history
		history := models.History{
			ChatId:      chatId,
			U1ToU2:      u1tou2,
			MessageType: utils.UserMessage,
			MessageBody: message,
			SendTime:    sendTime,
		}

		_, err = o.Insert(&history)
		if err != nil {
			fmt.Println("sendmessage insert history err: ", err)
			return
		}
	}

	fmt.Println("wow we are updating chat")

	//更新chat中最后一条信息
	_, err := o.QueryTable("chat").Filter("id", chatId).Update(orm.Params{
		"last_message": message,
		"u1_to_u2":     u1tou2,
		"unread_num":   orm.ColValue(orm.ColAdd, 1),
		"last_time":    sendTime,
	})
	if err != nil {
		fmt.Println("send message updating chat err: ", err)
		return
	}
}
