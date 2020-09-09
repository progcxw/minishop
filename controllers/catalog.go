package controllers

import (
	"minishop/models"
	"minishop/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CatalogController struct {
	beego.Controller
}

type CurCategory struct {
	models.MinishopCategory
	SubCategoryList []models.MinishopCategory `json:"subCategoryList"`
}

type CateLogIndexRtnJson struct {
	CategoryList    []models.MinishopCategory `json:"categoryList"`
	CurrentCategory CurCategory               `json:"currentCategory"`
}

func (this *CatalogController) Catalog_Index() {

	categoryId := this.GetString("id")

	o := orm.NewOrm()

	var categories []models.MinishopCategory
	categorytable := new(models.MinishopCategory)
	o.QueryTable(categorytable).Filter("parent_id", 0).Limit(10).All(&categories)

	var currentCategory *models.MinishopCategory = nil

	if categoryId != "" {
		o.QueryTable(categorytable).Filter("id", categoryId).One(currentCategory)
	}

	if currentCategory == nil {
		currentCategory = &categories[0]
	}

	curCategory := new(CurCategory)

	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []models.MinishopCategory
		o.QueryTable(categorytable).Filter("parent_id", currentCategory.Id).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.MinishopCategory = *currentCategory
	}

	utils.ReturnHTTPSuccess(&this.Controller, CateLogIndexRtnJson{categories, *curCategory})
	this.ServeJSON()

}

type CateLogCurRtnJson struct {
	CurrentCategory CurCategory `json:"currentCategory"`
}

func (this *CatalogController) Catalog_Current() {

	categoryId := this.GetString("id")

	o := orm.NewOrm()
	categorytable := new(models.MinishopCategory)
	currentCategory := new(models.MinishopCategory)
	if categoryId != "" {
		o.QueryTable(categorytable).Filter("id", categoryId).One(currentCategory)
	}

	curCategory := new(CurCategory)
	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []models.MinishopCategory
		o.QueryTable(categorytable).Filter("parent_id", currentCategory.Id).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.MinishopCategory = *currentCategory
	}

	utils.ReturnHTTPSuccess(&this.Controller, CateLogCurRtnJson{*curCategory})
	this.ServeJSON()

}
