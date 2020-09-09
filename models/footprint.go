package models

import (
	"minishop/utils"

	"github.com/astaxie/beego/orm"
)

func AddFootprint(userId, goodsId int) {
	o := orm.NewOrm()
	if userId > 0 && goodsId > 0 {
		footprintval := MinishopFootprint{GoodsId: goodsId, UserId: userId, AddTime: utils.GetTimestamp()}
		o.Insert(&footprintval)
	}
}
