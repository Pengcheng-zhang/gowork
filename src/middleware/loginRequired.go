package middleware

import (
	"fmt"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"model"
	"biz"
)
type LoginRequired struct {

}

func (loginRequire LoginRequired) Call(c martini.Context, r render.Render, req *http.Request, session sessions.Session) {
	v := session.Get("yz_session_token")
	var user model.UserModel
	if v != nil {
		fmt.Println("login required session:",v.(string))
		user = biz.GetUserFromSession(v.(string))
		if user.Id > 0 {
			c.Next()
		}else{
			fmt.Println("login required user id:", user.Id)
			if (req.Method == "POST") {
				r.JSON(200, map[string]interface{}{"code": 90000, "message" : "登录后才能进行此操作"})
				return
			}
			r.Redirect("/signin")
			return
		}
	}else{
		if (req.Method == "POST") {
		
			r.JSON(200, map[string]interface{}{"code": 90000, "message" : "登录后才能进行此操作"})
			return
		}
		r.Redirect("/signin")
	}
}