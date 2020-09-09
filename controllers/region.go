package controllers

import (
	"moshopserver/models"
	"moshopserver/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegionController struct {
	beego.Controller
}

func (this *RegionController) Region_Info() {

	regionId := this.GetString("regionId")
	intregionid := utils.String2Int(regionId)

	o := orm.NewOrm()
	regiontable := new(models.MinishopRegion)
	var region models.MinishopRegion
	o.QueryTable(regiontable).Filter("id", intregionid).One(&region)

	utils.ReturnHTTPSuccess(&this.Controller, region)
	this.ServeJSON()

}

func (this *RegionController) Region_List() {

	parentId := this.GetString("parentId")
	intparentid := utils.String2Int(parentId)

	o := orm.NewOrm()
	regiontable := new(models.MinishopRegion)
	var regions []models.MinishopRegion
	o.QueryTable(regiontable).Filter("parent_id", intparentid).All(&regions)

	utils.ReturnHTTPSuccess(&this.Controller, regions)
	this.ServeJSON()
}
