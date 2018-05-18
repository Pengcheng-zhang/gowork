package biz

import (
	"errors"
	"fmt"
	"model"
)
//帖子管理中心
type AvatarBiz struct{

}

//获取所有图像
func(this *AvatarBiz) GetAllAvatar(status string) []model.AvatarModel{
	var avatarModel []model.AvatarModel
	err := GetDbInstance().Where("status = ?", status).Scan(&avatarModel).Error
	if err != nil {
		fmt.Println("get avatar fail:", err)
	}
	return avatarModel
}

//添加头像资源
func(this *AvatarBiz) AddAvatar(avatarUrl string) (string, error){
	var avatar model.AvatarModel
	err := GetDbInstance().Where("url = ?", avatarUrl).First(&avatar).Error
	if avatar.Id > 0 {
		return "图像资源已存在", err
	}
	avatar = model.AvatarModel{ Url: avatarUrl, Status: "A" }
	err = GetDbInstance().Create(&avatar).Error
	fmt.Printf("user Register: err: %v\n", err)
	if err == nil{
		return "", nil
	}
	return "添加图像资源失败",errors.New("添加图像资源失败")
}

//删除头像资源
func(this *AvatarBiz) DeleteAvatar(avatar model.AvatarModel) bool{
	updateValue := map[string]interface{}{"status": "C"}
	err := GetDbInstance().Model(&avatar).Updates(updateValue).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}