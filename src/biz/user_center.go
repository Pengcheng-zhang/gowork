package biz

import (
	"fmt"
	"model"
	"strings"
	"encoding/base64"
	"strconv"
)

type UCenter struct {

}
type sessionStruct struct {
	UserID int
	Username string
	Roles string
}

var currentUser model.User;
func CheckUserExist(email string)  model.User{
	var user model.User
	GetDbInstance().Where("email = ?", email).First(&user)
	return user
}

func (uCenter UCenter) Register(email, password string)  bool{
	var user model.User
	user = CheckUserExist(email)
	if user.Id > 0 {
		return false
	}
	user = model.User{ Email:email, Password:password }
	err := GetDbInstance().Create(&user).Error
	fmt.Printf("user Register: user_id: %s\n", user.Id)
	fmt.Printf("user Register: err: %v\n", err)
	if err == nil{
		return true
	}
	return false
}

func (uCenter UCenter) Login(email string, password string) (string, model.User, error){
	var user model.User
	err := GetDbInstance().Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return "", user, err
	}
	fmt.Printf("user login: user_id: %s\n", user.Id)
	var originData string
	originData = strings.Join([]string{"user_id:", strconv.Itoa(user.Id), ";username:", user.Username, ";roles:", user.Roles},"")
	fmt.Printf("user login: origin data: %s\n", originData)
	result, err := AesEncrypt([]byte(originData))
	if err != nil {
		return "", user, err
	}
	var sessionString string
	sessionString = base64.StdEncoding.EncodeToString(result)
	fmt.Printf("user login: session: %s\n", sessionString)
	return sessionString, user,nil
}

func Logout() bool {
	return true
}

func GetCurrentUser() model.User {
	return currentUser
}

func setCurrentUser(user model.User) {
	currentUser = user
}

func GetUserFromSession(session string) model.User{
	var user model.User
	decodedSession,err := base64.StdEncoding.DecodeString(session)
	if err != nil {
		return user
	}
	originData, err := AesDecrypt(decodedSession)
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


