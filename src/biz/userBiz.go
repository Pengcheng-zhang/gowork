package biz

import (
	"model"
	"strings"
	"encoding/base64"
	"strconv"
	"regexp"
	"errors"
	"services"
)

type UserBiz struct {

}

type sessionStruct struct {
	UserID int
	Username string
	Roles string
}

var currentUser model.UserModel;
//检查用户是否存在
func CheckUserExist(username, email string)  (string,error){
	result := checkUserByName(username)
	if result == false {
		return "用户名已被占用", errors.New("username is used")
	}
	result = checkUserByEmail(email) 
	if result == false {
		return "邮箱已被占用", errors.New("email is used")
	}
	return "", nil
}

//检查用户名是否注册
func checkUserByName(username string) bool {
	var user model.UserModel
	GetDbInstance().Where("username = ?", username).First(&user)
	if user.Id > 0 {
		return false
	}
	return true
}

//检查邮箱是否注册
func checkUserByEmail (email string) bool {
	var user model.UserModel
	GetDbInstance().Where("email = ?", email).First(&user)
	if user.Id > 0 {
		return false
	}
	return true
}
//更新用户邮箱认证状态
func (this *UserBiz) UpdateUserVerifyStatus(email,status string) error{
	var user model.UserModel
	user.Email = email
   updateData := map[string]interface{}{"verified": status}
   _,err := this.UpdateUserInfo(user, updateData)
   return err
}
//检查敏感词
func (this *UserBiz) CheckSensitiveWord(username string) bool {
	r,_ := regexp.Compile("[*|/|\\|(|)]")
	result := r.FindStringIndex(username)
	if result != nil {
		return false
	}
	return true
}
//注册用户
func (this *UserBiz) Register(username, email, password string)  (string, bool){
	var user model.UserModel
	message, err := CheckUserExist(username, email)
	if err != nil {
		return message, false
	}
	user = model.UserModel{ Username: username, Email:email, Password:password }
	err = GetDbInstance().Create(&user).Error
	Debug("user Register: err:", err.Error())
	if err == nil{
		return "", true
	}
	return "注册失败",false
}
//用户登陆
func (this *UserBiz) Login(email string, password string) (string, model.UserModel, error){
	var user model.UserModel
	err := GetDbInstance().Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return "", user, err
	}
	var originData string
	originData = strings.Join([]string{"user_id:", strconv.Itoa(user.Id), ";username:", user.Username, ";roles:", user.Roles},"")
	Debug("user login: origin data:", originData)
	result, err := services.AesEncrypt([]byte(originData))
	if err != nil {
		Error("user login encrypt failed:", err.Error())
		return "", user, err
	}
	var sessionString string
	sessionString = base64.StdEncoding.EncodeToString(result)
	return sessionString, user,nil
}

//用户登出
func (this *UserBiz) Logout() bool {
	return true
}
//获取当前用户
func (this *UserBiz) GetCurrentUser() model.UserModel {
	return currentUser
}
//设置当前用户
func setCurrentUser(user model.UserModel) {
	currentUser = user
}
//从session中获取当前用户
func GetUserFromSession(session string) model.UserModel{
	var user model.UserModel
	if session == "" || len(session) == 0{
		return user
	}
	decodedSession,err := base64.StdEncoding.DecodeString(session)
	if err != nil {
		return user
	}
	originData, err := services.AesDecrypt(decodedSession)
	if err != nil {
		return user
	}
	var sessionString []string
	sessionString = strings.Split(string(originData),";")
	if sessionString == nil{
		return user
	}
	userIDString := strings.Split(sessionString[0], ":")
	if userIDString != nil && userIDString[0] == "user_id" {
		var userId int
		userId,err = strconv.Atoi(userIDString[1])
		if err != nil{
			return user
		}
		GetDbInstance().First(&user, userId)
		if user.Id > 0 {
			setCurrentUser(user)
		}
	}
	return user
}
//检查用户旧密码是否有效
func (this *UserBiz) CheckUserOldPassword(username string, oldPassword string) error{
	var user model.UserModel
	err := GetDbInstance().Where("username = ? AND password = ?", username, oldPassword).First(&user).Error
	return err
}
//更改用户信息
func (this *UserBiz) UpdateUserInfo(user model.UserModel, value interface{}) (model.UserModel,error) {
	err := GetDbInstance().Model(&user).Updates(value).Error
	if err != nil {
		Error("update user info failed:", err.Error())
	}
	return user,err
}
//签到
func (this *UserBiz) CheckIn(checkmodel model.SignHistoryModel) bool{
	err := GetDbInstance().Create(&checkmodel).Error
	if err != nil {
		Error("check in err:", err.Error())
		return false
	}
	return true
}
//检查今日是否签到
func (this *UserBiz) CheckedIn(checkmodel model.SignHistoryModel) bool{
	err := GetDbInstance().Where("to_days(created_at) = to_days(now())").First(&checkmodel).Error
	if err != nil {
		Error("check in today err:", err.Error())
		return false
	}
	return true
}
// 用户列表
func (this *UserBiz) GetUserList(limit int, offset int, status string) ([]model.UserModel, int){
	var userList []model.UserModel
	var err error
	count := 0
	DB := GetDbInstance()
	if status == "all" {
		DB.Find(&userList).Count(&count)
		err = DB.Limit(limit).Offset(offset).Find(&userList).Error
	} else {
		DB.Where("status = ?", status).Find(&userList).Count(&count)
		err = DB.Where("status = ?", status).Limit(limit).Offset(offset).Find(&userList).Error
	}
	if err != nil {
		Debug("get user list failed:", err.Error())
	}
	return userList, count
}


