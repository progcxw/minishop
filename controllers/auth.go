package controllers

import (
	"encoding/json"
	"fmt"

	"minishop/models"
	"minishop/services"
	"minishop/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AuthController struct {
	beego.Controller
}

type AuthLoginBody struct {
	Code     string               `json:"code"`
	UserInfo services.ResUserInfo `json:"userInfo"`
}

func (this *AuthController) Auth_LoginByWeixin() {
	var alb AuthLoginBody
	body := this.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &alb)
	clientIP := this.Ctx.Input.IP()
	userInfo := services.Login(alb.Code, alb.UserInfo)
	if userInfo == nil {
		fmt.Println("userInfo is nil")
		return
	}

	o := orm.NewOrm()

	var user models.MinishopUser
	usertable := new(models.MinishopUser)
	err = o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	if err == orm.ErrNoRows {
		newuser := models.MinishopUser{Username: utils.GetUUID(), Password: "", RegisterTime: utils.GetTimestamp(),
			RegisterIp: clientIP, Mobile: "", WeixinOpenid: userInfo.OpenID, Avatar: userInfo.AvatarUrl, Gender: userInfo.Gender,
			Nickname: userInfo.NickName}
		o.Insert(&newuser)
		o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	}

	userinfo := make(map[string]interface{})
	userinfo["id"] = user.Id
	userinfo["username"] = user.Username
	userinfo["nickname"] = user.Nickname
	userinfo["openId"] = user.WeixinOpenid
	userinfo["gender"] = user.Gender
	userinfo["avatar"] = user.Avatar
	userinfo["birthday"] = user.Birthday

	user.LastLoginIp = clientIP
	user.LastLoginTime = utils.GetTimestamp()

	if _, err := o.Update(&user); err == nil {

	}

	sessionKey := services.Create(utils.Int2String(user.Id))
	services.LoginUserId = utils.Int2String(user.Id)
	fmt.Println(services.LoginUserId)
	//fmt.Println("sessionkey==" + sessionKey)

	rtnInfo := make(map[string]interface{})
	rtnInfo["token"] = sessionKey
	rtnInfo["userInfo"] = userinfo

	utils.ReturnHTTPSuccess(&this.Controller, rtnInfo)
	this.ServeJSON()

}
