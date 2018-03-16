package main

import (
	"fmt"
	"time"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"biz"
	"route"
	//"yztest"
)

func main()  {
	now,err := time.Parse("2006-01-02 15:04:05 +0000 UTC", "2018-01-09 11:02:51 +0000 UTC" )
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	biz.DbInit();
	defer biz.GetDbInstance().Close()
	//var emailBiz biz.EmailBiz
	//emailBiz.UpdateUserVerifyStatus("zhangpch666@163.com", "Y")
	//emailBiz.Verify("zcqTGPGD8LvpseGUBjXXqOgOBLoX+Zlwj2eRMhmV+bMdL4cm2EPv+2u9u6J4mJ")
	//emailemailBizManage.CheckValid("770651352@qq.com")
	//return
	//yztest.Run()
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:"templates",
		Layout: "layout",
		Extensions:[]string{".tmpl",".html"},
		Charset:"UTF-8",
		IndentJSON: true,
	}))
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	route.Run(m)
}
