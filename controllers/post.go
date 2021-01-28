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

type PostController struct {
	beego.Controller
}

type locationMessage struct {
	City    string
	Message string
}

type postGoodObj struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Price   string `json:"price"`
	Postage string `json:"postage"`
	Images  []string
	CateId  int `json:"cateId"`
}

func (this *PostController) Post_CityLocation() {
	var location locationMessage
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body, &location)
	services.UserLocation = location.City
}

func (this *PostController) PostGoodsHandler() {
	var postGood postGoodObj
	var url string
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body, &postGood)

	for k, v := range postGood.Images {
		if k == 0 {
			url += v
		} else {
			url += "," + v
		}
	}

	o := orm.NewOrm()
	uploadGood := models.MinishopPostGoods{
		Name:        postGood.Name,
		Description: postGood.Desc,
		Price:       postGood.Price,
		Postage:     postGood.Postage,
		ListPicUrl:  url,
		CateId:      postGood.CateId,
		AddTime:     utils.GetTimestamp(),
		City:        services.UserLocation,
	}

	_, err := o.Insert(&uploadGood)
	if err != nil {
		fmt.Println("post insert error: ", err)
		return
	}
}
