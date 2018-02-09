package yztest

import (
	"fmt"
	"biz"
	"model"
)

var userManager biz.UCenterManager

func Run()  {
	checkIn()
}

//签到测试
func checkIn()  {
	var checkmodel model.SignHistoryModel
	checkmodel.UserId = 2
	result := userManager.CheckedIn(checkmodel)
	fmt.Println(result)
	if result == false {
		userManager.CheckIn(checkmodel)
	}
}