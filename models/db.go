package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

type MinishopAddress struct {
	Address    string `json:"address"`
	CityId     int    `json:"city_id"`
	CountryId  int    `json:"country_id"`
	DistrictId int    `json:"district_id"`
	Id         int    `json:"id"`
	IsDefault  int    `json:"is_default"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	ProvinceId int    `json:"province_id"`
	UserId     int    `json:"user_id"`
}

type MinishopAdmin struct {
	AddTime       int    `json:"add_time"`
	AdminRoleId   int    `json:"admin_role_id"`
	Avatar        string `json:"avatar"`
	Id            int    `json:"id"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime int    `json:"last_login_time"`
	Password      string `json:"password"`
	PasswordSalt  string `json:"password_salt"`
	UpdateTime    int    `json:"update_time"`
	Username      string `json:"username"`
}

type MinishopAttribute struct {
	AttributeCategoryId int    `json:"attribute_category_id"`
	Id                  int    `json:"id"`
	InputType           int    `json:"input_type"`
	Name                string `json:"name"`
	SortOrder           int    `json:"sort_order"`
	Values              string `json:"values"`
}

type MinishopAttributeCategory struct {
	Enabled int    `json:"enabled"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

type MinishopCart struct {
	Checked                   int     `json:"checked"`
	GoodsId                   int     `json:"goods_id"`
	GoodsName                 string  `json:"goods_name"`
	GoodsSn                   string  `json:"goods_sn"`
	GoodsSpecifitionIds       string  `json:"goods_specifition_ids"`
	GoodsSpecifitionNameValue string  `json:"goods_specifition_name_value"`
	Id                        int     `json:"id"`
	ListPicUrl                string  `json:"list_pic_url"`
	MarketPrice               float64 `json:"market_price"`
	Number                    int     `json:"number"`
	ProductId                 int     `json:"product_id"`
	RetailPrice               float64 `json:"retail_price"`
	SessionId                 string  `json:"session_id"`
	UserId                    int     `json:"user_id"`
}

type MinishopCategory struct {
	BannerUrl    string `json:"banner_url"`
	FrontDesc    string `json:"front_desc"`
	FrontName    string `json:"front_name"`
	IconUrl      string `json:"icon_url"`
	Id           int    `json:"id"`
	ImgUrl       string `json:"img_url"`
	IsShow       int    `json:"is_show"`
	Keywords     string `json:"keywords"`
	Level        string `json:"level"`
	Name         string `json:"name"`
	ParentId     int    `json:"parent_id"`
	ShowIndex    int    `json:"show_index"`
	SortOrder    int    `json:"sort_order"`
	Type         int    `json:"type"`
	WapBannerUrl string `json:"wap_banner_url"`
}

type MinishopChannel struct {
	IconUrl   string `json:"icon_url"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type MinishopCollect struct {
	AddTime     int64 `json:"add_time"`
	Id          int   `json:"id"`
	IsAttention int   `json:"is_attention"`
	TypeId      int   `json:"type_id"`
	UserId      int   `json:"user_id"`
	ValueId     int   `json:"value_id"`
}

type MinishopComment struct {
	AddTime    int64  `json:"add_time"`
	Content    string `json:"content"`
	Id         int    `json:"id"`
	NewContent string `json:"new_content"`
	Status     int    `json:"status"`
	TypeId     int    `json:"type_id"`
	UserId     int    `json:"user_id"`
	ValueId    int    `json:"value_id"`
}

type MinishopCommentPicture struct {
	CommentId int    `json:"comment_id"`
	Id        int    `json:"id"`
	PicUrl    string `json:"pic_url"`
	SortOrder int    `json:"sort_order"`
}

type MinishopFeedback struct {
	MessageImg string `json:"message_img"`
	MsgArea    int    `json:"msg_area"`
	MsgContent string `json:"msg_content"`
	Id         int    `json:"msg_id"`
	MsgStatus  int    `json:"msg_status"`
	MsgTime    int    `json:"msg_time"`
	MsgTitle   string `json:"msg_title"`
	MsgType    int    `json:"msg_type"`
	OrderId    int    `json:"order_id"`
	ParentId   int    `json:"parent_id"`
	UserEmail  string `json:"user_email"`
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
}

type MinishopFootprint struct {
	AddTime int64 `json:"add_time"`
	GoodsId int   `json:"goods_id"`
	Id      int   `json:"id"`
	UserId  int   `json:"user_id"`
}

type MinishopGoods struct {
	AddTime           int    `json:"add_time"`
	AppExclusivePrice string `json:"app_exclusive_price"`
	AttributeCategory int    `json:"attribute_category"`
	BrandId           int    `json:"brand_id"`
	CategoryId        int    `json:"category_id"`
	CounterPrice      string `json:"counter_price"`
	ExtraPrice        string `json:"extra_price"`
	GoodsBrief        string `json:"goods_brief"`
	GoodsDesc         string `json:"goods_desc"`
	GoodsNumber       int    `json:"goods_number"`
	GoodsSn           string `json:"goods_sn"`
	GoodsUnit         string `json:"goods_unit"`
	Id                int    `json:"id"`
	IsAppExclusive    int    `json:"is_app_exclusive"`
	IsDelete          bool   `json:"is_delete"`
	IsHot             int    `json:"is_hot"`
	IsLimited         int    `json:"is_limited"`
	IsNew             int    `json:"is_new"`
	IsOnSale          int    `json:"is_on_sale"`
	Keywords          string `json:"keywords"`
	ListPicUrl        string `json:"list_pic_url"`
	Name              string `json:"name"`
	PrimaryPicUrl     string `json:"primary_pic_url"`
	PrimaryProductId  int    `json:"primary_product_id"`
	PromotionDesc     string `json:"promotion_desc"`
	PromotionTag      string `json:"promotion_tag"`
	RetailPrice       string `json:"retail_price"`
	SellVolume        int    `json:"sell_volume"`
	SortOrder         int    `json:"sort_order"`
	UnitPrice         string `json:"unit_price"`
}

type MinishopPostGoods struct {
	Id          int    `json:"id"`
	AddTime     int64  `json:"add_time"`
	CateId      int    `json:"category_id"`
	Price       string `json:"price"`
	Postage     string `json:"postage"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	ListPicUrl  string `json:"list_pic_url"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Company     string `json:"company"`
}

type MinishopGoodsAttribute struct {
	AttributeId int    `json:"attribute_id"`
	GoodsId     int    `json:"goods_id"`
	Id          int    `json:"id"`
	Value       string `json:"value"`
}

type MinishopGoodsGallery struct {
	GoodsId   int    `json:"goods_id"`
	Id        int    `json:"id"`
	ImgDesc   string `json:"img_desc"`
	ImgUrl    string `json:"img_url"`
	SortOrder int    `json:"sort_order"`
}

type MinishopGoodsIssue struct {
	Answer   string `json:"answer"`
	GoodsId  string `json:"goods_id"`
	Id       int    `json:"id"`
	Question string `json:"question"`
}

type MinishopGoodsSpecification struct {
	GoodsId         int    `json:"goods_id"`
	Id              int    `json:"id"`
	PicUrl          string `json:"pic_url"`
	SpecificationId int    `json:"specification_id"`
	Value           string `json:"value"`
}

type MinishopKeywords struct {
	Id        int    `json:"id"`
	IsDefault int    `json:"is_default"`
	IsHot     int    `json:"is_hot"`
	IsShow    int    `json:"is_show"`
	Keyword   string `json:"keyword"`
	SchemeUrl string `json:"scheme_url"`
	SortOrder int    `json:"sort_order"`
	Type      int    `json:"type"`
}

type MinishopOrder struct {
	ActualPrice    float64 `json:"actual_price"`
	AddTime        int64   `json:"add_time"`
	Address        string  `json:"address"`
	CallbackStatus string  `json:"callback_status"`
	City           int     `json:"city"`
	ConfirmTime    int     `json:"confirm_time"`
	Consignee      string  `json:"consignee"`
	Country        int     `json:"country"`
	CouponId       int     `json:"coupon_id"`
	CouponPrice    float64 `json:"coupon_price"`
	District       int     `json:"district"`
	FreightPrice   float64 `json:"freight_price"`
	GoodsPrice     float64 `json:"goods_price"`
	Id             int     `json:"id"`
	Integral       int     `json:"integral"`
	IntegralMoney  float64 `json:"integral_money"`
	Mobile         string  `json:"mobile"`
	OrderPrice     float64 `json:"order_price"`
	OrderSn        string  `json:"order_sn"`
	OrderStatus    int     `json:"order_status"`
	ParentId       int     `json:"parent_id"`
	PayId          int     `json:"pay_id"`
	PayName        string  `json:"pay_name"`
	PayStatus      int     `json:"pay_status"`
	PayTime        int     `json:"pay_time"`
	Postscript     string  `json:"postscript"`
	Province       int     `json:"province"`
	ShippingFee    float64 `json:"shipping_fee"`
	ShippingStatus int     `json:"shipping_status"`
	UserId         int     `json:"user_id"`
}

type MinishopOrderExpress struct {
	AddTime      int    `json:"add_time"`
	Id           int    `json:"id"`
	IsFinish     int    `json:"is_finish"`
	LogisticCode string `json:"logistic_code"`
	OrderId      int    `json:"order_id"`
	RequestCount int    `json:"request_count"`
	RequestTime  int    `json:"request_time"`
	ShipperCode  string `json:"shipper_code"`
	ShipperId    int    `json:"shipper_id"`
	ShipperName  string `json:"shipper_name"`
	Traces       string `json:"traces"`
	UpdateTime   int    `json:"update_time"`
}

type MinishopOrderGoods struct {
	GoodsId                   int     `json:"goods_id"`
	GoodsName                 string  `json:"goods_name"`
	GoodsSn                   string  `json:"goods_sn"`
	GoodsSpecifitionIds       string  `json:"goods_specifition_ids"`
	GoodsSpecifitionNameValue string  `json:"goods_specifition_name_value"`
	Id                        int     `json:"id"`
	IsReal                    int     `json:"is_real"`
	ListPicUrl                string  `json:"list_pic_url"`
	MarketPrice               float64 `json:"market_price"`
	Number                    int     `json:"number"`
	OrderId                   int     `json:"order_id"`
	ProductId                 int     `json:"product_id"`
	RetailPrice               float64 `json:"retail_price"`
}

type MinishopProduct struct {
	GoodsId               int     `json:"goods_id"`
	GoodsNumber           int     `json:"goods_number"`
	GoodsSn               string  `json:"goods_sn"`
	GoodsSpecificationIds string  `json:"goods_specification_ids"`
	Id                    int     `json:"id"`
	RetailPrice           float64 `json:"retail_price"`
}

type MinishopRegion struct {
	AgencyId int    `json:"agency_id"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
	Type     int    `json:"type"`
}

type MinishopRelatedGoods struct {
	GoodsId        int `json:"goods_id"`
	Id             int `json:"id"`
	RelatedGoodsId int `json:"related_goods_id"`
}

type MinishopSearchHistory struct {
	AddTime int64  `json:"add_time"`
	From    string `json:"from"`
	Id      int    `json:"id"`
	Keyword string `json:"keyword"`
	UserId  string `json:"user_id"`
}

type MinishopShipper struct {
	Code      string `json:"code"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type MinishopSpecification struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type MinishopUser struct {
	Avatar        string `json:"avatar"`
	Birthday      int    `json:"birthday"`
	Gender        int    `json:"gender"`
	Id            int    `json:"id"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime int64  `json:"last_login_time"`
	Mobile        string `json:"mobile"`
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	RegisterIp    string `json:"register_ip"`
	RegisterTime  int64  `json:"register_time"`
	UserLevelId   int    `json:"user_level_id"`
	Username      string `json:"username"`
	WeixinOpenid  string `json:"weixin_openid"`
}

type MinishopUserLevel struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
}

//MinishopPic : 图片储存结构体
type MinishopPic struct {
	Id       int
	FileSha1 string
	FileName string
	FileSize int64
	FileAddr string
}

type Chat struct {
	Id          int64  `json:"id"`
	U1          int    `json:"u1"`
	U2          int    `json:"u2"`
	GoodsId     int    `json:"goods_id"`
	ShowToU1    bool   `json:"show_to_u1"`
	ShowToU2    bool   `json:"show_to_u2"`
	LastMessage string `json:"last_message"`
	U1ToU2      bool   `json:"u1_to_u2"`
	UnreadNum   int    `json:"unread_num"`
	LastTime    int64  `json:"last_time"`
}

type History struct {
	Id          int64  `json:"id"`
	ChatId      int64  `json:"chat_id"`
	U1ToU2      bool   `json:"u1_to_u2"`
	MessageType int    `json:"message_type"`
	MessageBody string `json:"message_body"`
	SendTime    int64  `json:"send_time"`
}

// type User struct {
// 	Id   int    `orm:"not null pk autoincr INT(11)"`
// 	Name string `orm:"not null default '' VARCHAR(100)"`
// }

func init() {

	// set default database
	err := orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1:3306)/minishop", 30)
	if err != nil {
		fmt.Println(err)
	}

	// register model
	orm.RegisterModel(new(MinishopAddress))
	orm.RegisterModel(new(MinishopAdmin))
	orm.RegisterModel(new(MinishopAttribute))
	orm.RegisterModel(new(MinishopAttributeCategory))

	orm.RegisterModel(new(MinishopCart))
	orm.RegisterModel(new(MinishopCategory))

	orm.RegisterModel(new(MinishopChannel))
	orm.RegisterModel(new(MinishopCollect))
	orm.RegisterModel(new(MinishopComment))
	orm.RegisterModel(new(MinishopCommentPicture))

	orm.RegisterModel(new(MinishopFeedback))

	orm.RegisterModel(new(MinishopFootprint))
	orm.RegisterModel(new(MinishopGoods))

	orm.RegisterModel(new(MinishopGoodsAttribute))

	orm.RegisterModel(new(MinishopGoodsGallery))
	orm.RegisterModel(new(MinishopGoodsIssue))

	orm.RegisterModel(new(MinishopGoodsSpecification))
	orm.RegisterModel(new(MinishopKeywords))

	orm.RegisterModel(new(MinishopOrder))
	orm.RegisterModel(new(MinishopOrderExpress))

	orm.RegisterModel(new(MinishopOrderGoods))

	orm.RegisterModel(new(MinishopProduct))
	orm.RegisterModel(new(MinishopRegion))

	orm.RegisterModel(new(MinishopRelatedGoods))
	orm.RegisterModel(new(MinishopSearchHistory))

	orm.RegisterModel(new(MinishopShipper))
	orm.RegisterModel(new(MinishopSpecification))

	orm.RegisterModel(new(MinishopUser))
	orm.RegisterModel(new(MinishopUserLevel))

	orm.RegisterModel(new(MinishopPic))
	orm.RegisterModel(new(MinishopPostGoods))

	orm.RegisterModel(new(Chat))
	orm.RegisterModel(new(History))
}
