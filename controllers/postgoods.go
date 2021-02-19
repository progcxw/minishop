package controllers

import (
	"encoding/base64"
	"minishop/models"
	"minishop/utils"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PostGoodsController struct {
	beego.Controller
}

type postGoodsRtnJson struct {
	Goods          models.MinishopPostGoods `json:"info"`
	Galleries      []string                 `json:"gallery"`
	Seller         models.MinishopUser      `json:"seller"`
	UserHasCollect int                      `json:"userHasCollect"`
	Comment        Comment                  `json:"comment"`
}

func (this *PostGoodsController) Postgoods_Detail() {
	goodsId := this.GetString("id")
	intGoodsId := utils.String2Int(goodsId)

	o := orm.NewOrm()

	var goodone models.MinishopPostGoods
	o.QueryTable("minishop_post_goods").Filter("id", intGoodsId).One(&goodone)

	var seller models.MinishopUser
	o.QueryTable("minishop_user").Filter("id", goodone.UserId).One(&seller)

	comment := new(models.MinishopComment)
	commentCount, _ := o.QueryTable(comment).Filter("value_id", intGoodsId).Filter("type_id", 0).Count()
	var hotcommentone models.MinishopComment
	o.QueryTable(comment).Filter("value_id", intGoodsId).Filter("type_id", 0).One(&hotcommentone)

	var commentInfo CommentInfo

	if &hotcommentone != nil {
		user := new(models.MinishopUser)
		var commentUsers []orm.Params
		o.QueryTable(user).Filter("id", hotcommentone.UserId).Values(&commentUsers, "nickname", "username", "avatar")
		content, _ := base64.StdEncoding.DecodeString(hotcommentone.Content)

		var commentpictures []models.MinishopCommentPicture
		commentpicture := new(models.MinishopCommentPicture)
		o.QueryTable(commentpicture).Filter("comment_id", hotcommentone.Id).All(&commentpictures)

		commentInfo = CommentInfo{Content: string(content), AddTime: hotcommentone.AddTime, NickName: user.Nickname, Avatar: user.Avatar, PicList: commentpictures}
	}

	commentval := Comment{Count: commentCount, Data: commentInfo}
	loginuserid := getLoginUserId()

	userhascollect := models.IsUserHasCollect(loginuserid, 0, intGoodsId)

	models.AddFootprint(loginuserid, intGoodsId)

	picList := strings.Split(goodone.ListPicUrl, ",")

	utils.ReturnHTTPSuccess(&this.Controller, postGoodsRtnJson{
		Goods:          goodone,
		Galleries:      picList,
		Seller:         seller,
		UserHasCollect: userhascollect,
		Comment:        commentval,
	})

	this.ServeJSON()
}
