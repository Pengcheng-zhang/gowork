package biz

import (
	"model"
)

type FriendsBiz struct{

}

func(this *FriendsBiz) GetList(sex int, limit int, offset int) []model.FriendsModel{
	var friends []model.FriendsModel
	err := GetDbInstance().Where("sex = ?", sex).Limit(limit).Offset(offset).Find(&friends).Error
	if err != nil {
		Debug("get friends list failed:", err.Error())
	}
	return friends
}

func(this *FriendsBiz) Detail(id int) model.FriendsModel {
	var friend model.FriendsModel
	err := GetDbInstance().Where("id = ?", id).First(&friend).Error
	if err != nil {
		Debug("get friend failed:", err.Error())
	}else{
		err = GetDbInstance().Model(&friend).Update("view_times", friend.ViewTimes + 1).Error
		if err != nil {
			Debug("update friend view times failed:", err.Error())
		}
	}
	return friend
}

func(this *FriendsBiz) Create(friend model.FriendsModel) int{
	err := GetDbInstance().Create(&friend).Error
	if err != nil {
		Debug("friend create error:%s", err.Error())
		return 0
	}
	Debug("friend id is : %d", friend.Id)
	return friend.Id
}

func (this *FriendsBiz) Update(friend model.FriendsModel, value interface{}) bool{
	err := GetDbInstance().Model(&friend).Updates(value).Error
	if err != nil {
		Debug("update friend failed:", err.Error())
		return false
	}
	return true
}