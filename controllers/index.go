package controllers

import (
	"minishop/models"
	"minishop/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

type newCategoryList struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	GoodsList []orm.Params `json:"goodsList"`
}

type IndexRtnJson struct {
	Channels     []models.MinishopChannel `json:"channel"`
	CategoryList []newCategoryList        `json:"categoryList"`
}

func updateJsonKeysIndex(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "ListPicUrl":
				delete(val, k)
				val["list_pic_url"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			}
		}
	}
}

func (this *IndexController) Index_Index() {
	o := orm.NewOrm()

	var channels []models.MinishopChannel
	channel := new(models.MinishopChannel)
	o.QueryTable(channel).OrderBy("sort_order").All(&channels)

	var newgoods []orm.Params
	goods := new(models.MinishopGoods)
	o.QueryTable(goods).Filter("is_new", 1).Limit(4).Values(&newgoods, "id", "name", "list_pic_url", "retail_price")
	updateJsonKeysIndex(newgoods)

	var categoryList []models.MinishopCategory
	category := new(models.MinishopCategory)
	o.QueryTable(category).Filter("parent_id", 0).Exclude("name", "推荐").All(&categoryList)

	var newList []newCategoryList

	for _, categoryItem := range categoryList {
		var mapids []orm.Params
		o.QueryTable(category).Filter("parent_id", categoryItem.Id).Values(&mapids, "id")

		// var valIds []int64
		// for _, value := range mapids {
		// 	valIds = append(valIds, value["Id"].(int64))
		// }

		valIds := utils.ExactMapValues2Int64Array(mapids, "Id")

		var categorygoods []orm.Params
		o.QueryTable(goods).Filter("category_id__in", valIds).Limit(7).Values(&categorygoods, "id", "name", "list_pic_url", "retail_price")
		updateJsonKeysIndex(categorygoods)
		newList = append(newList, newCategoryList{categoryItem.Id, categoryItem.Name, categorygoods})
	}

	utils.ReturnHTTPSuccess(&this.Controller, IndexRtnJson{channels, newList})

	this.ServeJSON()

}
