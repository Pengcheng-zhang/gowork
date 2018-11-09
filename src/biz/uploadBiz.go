package biz

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"common"
)

type UploadBiz struct{
	ossHost string
}

func(this *UploadBiz) getBucket() (*oss.Bucket){
	ossConfig := common.GetConfigSection("oss")
	if ossConfig == nil {
		Debug("get oss config failed")
		return nil
	}
	this.ossHost = "https://"+ ossConfig["bucketName"] + "." + ossConfig["endpoint"] + "/"
	// 创建OSSClinet实例
	client, err := oss.New(ossConfig["endpoint"], ossConfig["accessKeyId"], ossConfig["accessKeySecret"])
	if err != nil {
		Debug("Create oss client failed:", err.Error())
		return nil
	}
	// 获取存储空间
	bucket, err := client.Bucket(ossConfig["bucketName"])
	if err != nil {
		Debug("Get oss bucket failed:", err.Error())
		return nil
	}
	return bucket
}

func(this *UploadBiz) Upload (objectName string, file multipart.File) (result bool, url string){
	bucket := this.getBucket()
	//上传文件
	options := []oss.Option{
		oss.ContentType("image/jpeg"),
	}
	
	err := bucket.PutObject(objectName, file, options...)
	if err != nil {
		Debug("Put object from file failed:", err.Error())
		return false, ""
	}
	return true, this.ossHost+objectName
}

func(this *UploadBiz) Delete(objectName string) bool{
	// 获取存储空间。
	bucket := this.getBucket()
	// 删除单个文件。
	err := bucket.DeleteObject(objectName)
	if err != nil {
		Debug("Delete object from file failed:", err.Error())
		return false
	}
	return true
}