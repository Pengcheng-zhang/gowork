package controller

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"strconv"
	"biz"
)

type FriendsController struct {
	fBiz biz.FriendsBiz
}

func(this *FriendsController) List(r render.Render, req *http.Request, session sessions.Session)  {
	sex := req.FormValue("sex")
	limit := req.FormValue("limit")
	offset := req.FormValue("offset")

	var querySex, queryLimit, queryOffset int
	querySex,err := strconv.Atoi(sex)
	if err != nil {
		querySex = 0
	}
	queryLimit,err = strconv.Atoi(limit)
	if err != nil {
		queryLimit = 10
	}
	queryOffset,err = strconv.Atoi(offset)
	if err != nil {
		queryOffset = 0
	}
	result := this.fBiz.GetList(querySex, queryLimit, queryOffset)
	SetCommonResult(10000, "success", result)
	r.JSON(200, CommonResult)
}

func(this *FriendsController) Detail(r render.Render, req *http.Request, session sessions.Session) {
	id := req.FormValue("id")
	queryId,err := strconv.Atoi(id)
	if err != nil {
		SetCommonResult(40004, "fail", "未找到你心仪的对象")
		r.JSON(200, CommonResult)
		return
	}
	result := this.fBiz.Detail(queryId)
	SetCommonResult(10000, "success", result)
	r.JSON(200, CommonResult)
}

func(this *FriendsController) New(r render.Render, req *http.Request, session sessions.Session) {

}

func(this *FriendsController) Update(r render.Render, req *http.Request, session sessions.Session) {
	
}