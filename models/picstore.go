package models

import (
	"github.com/astaxie/beego/orm"
)

// OnFileUploadFinished : 文件上传完成，保存meta
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	o := orm.NewOrm()
	pic := MinishopPic{
		FileHash: filehash,
		FileName: filename,
		FileSize: filesize,
		FileAddr: fileaddr,
	}

	_, err := o.Insert(&pic)
	if err != nil {
		return false
	}

	return true
}

// GetFileMeta : 从mysql获取文件元信息
func GetFileMeta(filehash string) MinishopPic {
	o := orm.NewOrm()
	var pics []MinishopPic
	pic := new(MinishopPic)
	o.QueryTable(pic).Filter("FileHash", filehash).One(&pics)

	return pics[0]
}

//DelFileMeta : 删除文件元信息
func DelFileMeta(filehash string) bool {
	o := orm.NewOrm()
	pic := new(MinishopPic)

	_, err := o.QueryTable(pic).Filter("FileHash", filehash).Delete()
	if err != nil {
		return false
	}

	return true
}
