package models

import (
	"minishop/utils"

	"github.com/astaxie/beego/orm"
)

func GetChildCategoryId(categoryid int) []int64 {

	o := orm.NewOrm()
	categorytable := new(MinishopCategory)
	var childids []orm.Params
	o.QueryTable(categorytable).Filter("parent_id", categoryid).Limit(10000).Values(&childids, "id")
	childintids := utils.ExactMapValues2Int64Array(childids, "Id")
	return childintids
}

func GetCategoryWhereIn(categoryid int) []int64 {

	childintids := GetChildCategoryId(categoryid)
	childintids = append(childintids, int64(categoryid))
	return childintids
}
