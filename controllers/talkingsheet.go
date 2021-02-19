package controllers

import (
	"minishop/models"
	"minishop/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SheetsController struct {
	beego.Controller
}

func (this *SheetsController) SheetsIndex() {
	o := orm.NewOrm()
	var sheets []models.Sheets
	o.QueryTable("sheets").All(&sheets)

	utils.ReturnHTTPSuccess(&this.Controller, sheets)
	this.ServeJSON()
}

func (this *SheetsController) SheetsMsg() {
	msg := this.GetString("message")
	o := orm.NewOrm()

	sheet := models.Sheets{
		Content: msg,
	}

	o.Insert(&sheet)
}
