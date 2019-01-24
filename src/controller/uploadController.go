package controller

import (
	"github.com/martini-contrib/render"
	"net/http"
	"encoding/base64"
	"time"
	"strconv"
	"path"
	"biz"
)

type UploadController struct {
	uploadBiz biz.UploadBiz
}

func(this *UploadController) File(r render.Render, req *http.Request) {
	req.ParseForm()
	file, handler, err := req.FormFile("loadfile")
	if err != nil {
		Debug("get upload file failed", err.Error())
		setJsonResult(40001, "请选择文件", "")
		r.JSON(200, jsonResult)
		return
	}
	defer file.Close()
	fileSuffix := path.Ext(handler.Filename)
	if fileSuffix != ".png" && fileSuffix != ".jpg" {
		setJsonResult(40002, "请选择图片文件", "")
		r.JSON(200, jsonResult)
		return
	}
	Debug("upload file size:", handler.Size)
	if handler.Size > 5000000 {
		setJsonResult(40003, "文件过大，请选择小于5M的文件", "")
		r.JSON(200, jsonResult)
		return
	}
	currentTime := time.Now().Unix()
	user_id := "123/"
	objectName := user_id + base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(currentTime, 10)))
	Debug("upload file name:", objectName)
	result, url := this.uploadBiz.Upload(objectName, file)
	if result {
		setJsonResult(40004, "上传成功", url)
		Debug("upload file url:", url)
		r.JSON(200, jsonResult)
	}else {
		setJsonResult(10000, "上传错误", "")
		r.JSON(200, jsonResult)
	}
	return
}
