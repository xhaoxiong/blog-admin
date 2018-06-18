package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"time"

	"github.com/astaxie/beego"
)

// @router /api/admin/file/upload [*]
func (this *AdminController) Upload() {
	_, fileHeader, err := this.GetFile("file")
	if err != nil {
		beego.Debug("get file error :", err)
	}
	result := make(map[string]interface{})
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		result["stauts"] = 10001
		result["message"] = "error"
		beego.Debug("Open file error :", err)

	}
	filename := fileHeader.Filename + string(time.Now().Unix())
	f, err := os.OpenFile("upload/file/"+filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	io.Copy(f, file)
	hashfile, err := os.Open("upload/file/" + filename)
	hash := md5.New()
	if _, err := io.Copy(hash, hashfile); err != nil {
		beego.Error("get hash error:", err)
		this.ReturnJson(10002, "get hash error")
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	hashString := hex.EncodeToString(hashInBytes)
	hashfile.Close()

	if err := os.Rename("upload/file/"+filename, "upload/file/"+hashString); err != nil {
		beego.Debug("file  rename error :", err)
	}

	result["status"] = 10000
	result["message"] = "success"
	result["filepath"] = "/upload/file/" + hashString
	this.Data["json"] = result
	this.ServeJSON()
	return

}
