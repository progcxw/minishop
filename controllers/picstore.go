package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"minishop/models"
	"minishop/utils"

	"github.com/astaxie/beego"
)

type PicstoreController struct {
	beego.Controller
}

type JsonType struct {
	code int
	msg  string
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
	defer func() {
		this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
		this.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if errCode < 0 {
			this.Data["json"] = &JsonType{errCode, "上传失败"}
			this.ServeJSON()
		} else {
			this.Data["json"] = &JsonType{errCode, "上传成功"}
			this.ServeJSON()
		}
	}()

	fmt.Println("going into file upload")
	// 1. 从form表单中获得文件内容句柄
	file, head, err := this.GetFile("file")
	if err != nil {
		log.Printf("Failed to get form data, err:%s\n", err.Error())
		errCode = -1
		return
	}
	defer file.Close()

	// 2. 把文件内容转为[]byte
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Printf("Failed to get file data, err:%s\n", err.Error())
		errCode = -2
		return
	}

	// 3. 构建文件元信息
	fileMeta := models.MinishopPic{
		FileName: head.Filename,
		FileHash: utils.Sha1(buf.Bytes()), //　计算文件sha1
		FileSize: int64(len(buf.Bytes())),
	}

	// 4. 将文件写入临时存储位置
	fileMeta.FileAddr = utils.MergeLocalRootDir + fileMeta.FileName // 存储地址
	newFile, err := os.Create(fileMeta.FileAddr)
	if err != nil {
		log.Printf("Failed to create file, err:%s\n", err.Error())
		errCode = -3
		return
	}
	defer newFile.Close()

	nByte, err := newFile.Write(buf.Bytes())
	if int64(nByte) != fileMeta.FileSize || err != nil {
		log.Printf("Failed to save data into file, writtenSize:%d, err:%s\n", nByte, err.Error())
		errCode = -4
		return
	}

	//6.  更新文件表记录
	ret := models.OnFileUploadFinished(fileMeta.FileHash, fileMeta.FileName, fileMeta.FileSize, fileMeta.FileAddr)
	if !ret {
		errCode = -6
		return
	}

	// 7. 更新用户文件表
}
