package middleware

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"model"
	"biz"
)
type LoginRequired struct {

}

func (loginRequire LoginRequired) Call(c martini.Context, r render.Render, session sessions.Session) {
	v := session.Get("sucai_session_token")
	var user model.User
	if v != nil {
		fmt.Println(v)
		user = biz.GetUserFromSession(v.(string))
		if user.Id > 0 {
			c.Next()
		}else{
			r.Redirect("/login")
		}
	}
	r.Redirect("/login")
}