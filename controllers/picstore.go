package controllers

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"minishop/models"
	"minishop/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PicstoreController struct {
	beego.Controller
}

type JsonType struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Url  string `json:"filesha1"`
}

func init() {
	if err := os.MkdirAll(utils.TempLocalRootDir, 0744); err != nil {
		fmt.Println("无法指定目录用于存储临时文件: " + utils.TempLocalRootDir)
		os.Exit(1)
	}
	if err := os.MkdirAll(utils.MergeLocalRootDir, 0744); err != nil {
		fmt.Println("无法指定目录用于存储合并后文件: " + utils.MergeLocalRootDir)
		os.Exit(1)
	}
}

// DoUploadHandler ： 处理文件上传
func (this *PicstoreController) DoUploadHandler() {
	errCode := 0
	var url string

	defer func() {
		if errCode < 0 {
			utils.ReturnHTTPSuccess(&this.Controller, JsonType{errCode, "上传失败", ""})
			this.ServeJSON()
		} else {
			utils.ReturnHTTPSuccess(&this.Controller, JsonType{errCode, "上传成功", url})
			this.ServeJSON()
		}
	}()

	// 1. 从form表单中获得文件内容句柄
	file, head, err := this.GetFile("file")
	if err != nil {
		fmt.Printf("Failed to get form data, err:%s\n", err.Error())
		errCode = -1
		return
	}
	defer file.Close()

	// 2. 把文件内容转为[]byte
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Printf("Failed to get file data, err:%s\n", err.Error())
		errCode = -2
		return
	}

	// 3. 构建文件元信息
	fileMeta := models.MinishopPic{
		FileName: head.Filename,
		FileSha1: utils.Sha1(buf.Bytes()), //　计算文件sha1
		FileSize: int64(len(buf.Bytes())),
	}

	url = utils.PicStoreAPI + fileMeta.FileSha1

	// 4. 将文件写入临时存储位置
	fileMeta.FileAddr = utils.MergeLocalRootDir + fileMeta.FileName // 存储地址
	newFile, err := os.Create(fileMeta.FileAddr)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		errCode = -3
		return
	}
	defer newFile.Close()

	nByte, err := newFile.Write(buf.Bytes())
	if int64(nByte) != fileMeta.FileSize || err != nil {
		fmt.Printf("Failed to save data into file, writtenSize:%d, err:%s\n", nByte, err.Error())
		errCode = -4
		return
	}

	//6.  更新文件表记录
	ret := models.OnFileUploadFinished(fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.FileAddr)
	if !ret {
		errCode = -6
		return
	}

	// 7. 更新用户文件表
}

// DownloadHandler : 文件下载接口
func (this *PicstoreController) DownloadHandler() {
	fileSha1 := this.GetString("fileSha1")

	o := orm.NewOrm()
	fileMeta := new(models.MinishopPic)
	o.QueryTable(fileMeta).Filter("fileSha1", fileSha1).One(fileMeta)

	this.Ctx.Output.Download(fileMeta.FileAddr, fileMeta.FileName)
}
