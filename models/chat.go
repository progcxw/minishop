package models

import "time"

//SimpleUser 储存聊天页面的用户信息
type SimpleUser struct {
	UserId    int    `json:"user_id"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}

//SimpleGoods 储存聊天页面的相关商品信息
type SimpleGoods struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	PrimaryPicUrl string `json:"primary_pic_url"`
	Price         string `json:"price"`
}

//ChatForm 聊天框
type ChatForm struct {
	OtherSide   SimpleUser  `json:"other_side"`
	Goods       SimpleGoods `json:"goods"`
	IsU1        bool        `json:"is_u1"`
	HistoryList []History   `json:"history_list"`
	OffsetTime  int64       `json:"offset_time"`
}

//ChatIndexEle 消息列表的各个消息
type ChatIndexEle struct {
	UnreadNum int         `json:"unread_num"`
	OtherSide SimpleUser  `json:"other_side"`
	Goods     SimpleGoods `json:"simple_goods"`
	LastChat  History     `json:"last_chat"`
	UserId    int         `json:"user_id"`
	GoodsId   int         `json:"goods_id"`
}

//ChatIndex 消息列表页面
type ChatIndex struct {
	Chats      []ChatIndexEle `json:"chats"`
	OffsetTime int64          `json:"offset_time"`
}

//WsMessage websocket消息
type WsMessage struct {
	ChatId      int64  `json:"chat_id"`
	SenderId    int    `json:"sender_id"`
	ReceiverId  int    `json:"receiver_id"`
	GoodsId     int    `json:"goods_id"`
	MessageType int    `json:"message_type"`
	MessageBody string `json:"message_body"`
	SendTime    int64  `json:"send_time"`
}

//ChatEleList for聊天组件的排序
type ChatEleList []ChatIndexEle

func (x ChatEleList) Len() int {
	return len(x)
}

func (x ChatEleList) Less(i, j int) bool {
	return time.Unix(x[i].LastChat.SendTime, 0).After(time.Unix(x[j].LastChat.SendTime, 0))
}

func (x ChatEleList) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func WsListToHisList(wsMessageList []WsMessage) []History {
	if wsMessageList == nil {
		return nil
	}

	var u1ToU2 bool
	historyList := make([]History, len(wsMessageList))

	for i, wsMsg := range wsMessageList {
		if wsMsg.SenderId > wsMsg.ReceiverId {
			u1ToU2 = false
		} else {
			u1ToU2 = true
		}

		historyList[i].U1ToU2 = u1ToU2
		historyList[i].MessageBody = wsMsg.MessageBody
		historyList[i].MessageType = wsMsg.MessageType
		historyList[i].SendTime = wsMsg.SendTime
	}

	return historyList
}
