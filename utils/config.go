package utils

import (
	"fmt"

	"github.com/astaxie/beego"
)

//PicStoreAPI 用于生成图片url
const PicStoreAPI = "http://81.71.134.143:8360/api/image/get?fileSha1="

// 存储类型(表示文件存到哪里)
type StoreType int

const (
	_ StoreType = iota
	// StoreLocal : 节点本地
	StoreLocal
	// StoreCeph : Ceph集群
	StoreCeph
	// StoreOSS : 阿里OSS
	StoreOSS
	// StoreMix : 混合(Ceph及OSS)
	StoreMix
	// StoreAll : 所有类型的存储都存一份数据
	StoreAll
)

const (
	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir = "/home/ubuntu/fileserver_tmp/"
	// MergeLocalRootDir : 本地存储地址的路径(包含普通上传及分块上传)
	MergeLocalRootDir = "/home/ubuntu/fileserver_merge/"
	// ChunckLocalRootDir : 分块存储地址的路径
	ChunckLocalRootDir = "/home/ubuntu/fileserver_chunk/"
	// CephRootDir : Ceph的存储路径prefix
	CephRootDir = "/home/ubuntu/ceph"
	// OSSRootDir : OSS的存储路径prefix
	OSSRootDir = "/home/ubuntu/oss/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = StoreLocal
)

const (
	//SysMessage 系统消息
	SysMessage = 0
	//UserMessage 用户发送的消息
	UserMessage = 1
	//EstablishChat 单方建立的对话，不展示给对方
	EstablishChat = 2
	//FirstChat 第一次给对方发送消息，双方可见
	FirstChat = 3
	//UnreadNum 未读消息
	UnreadNum = 4
)

/*
  default_module: 'api'
  weixin:
    appid: '' #小程序 appid
    secret: '' #小程序密钥
    mch_id: '' #商户帐号ID
    partner_key: '' #微信支付密钥
    notify_url: '' #微信异步通知，例：https://www.minishop.com/api/pay/notify
  express:
    #快递物流信息查询使用的是快递鸟接口，申请地址：http://www.kdniao.com/
    appid: ''  #对应快递鸟用户后台 用户ID
    appkey: '' #对应快递鸟用户后台 API key
    request_url: 'http://api.kdniao.cc/Ebusiness/EbusinessOrderHandle.aspx'
*/

func init() {
	err := beego.LoadAppConfig("ini", "conf/config.conf")
	if err != nil {
		fmt.Println("config load error: ", err)
	}
}
