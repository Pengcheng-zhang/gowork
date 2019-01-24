package middleware

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"strings"
	"model"
	"biz"
)
type AdminRequired struct {

}

func (adminRequired AdminRequired) Call(c martini.Context, r render.Render, req *http.Request, session sessions.Session) {
	v := session.Get("yz_session_token")
	var user model.UserModel
	if v != nil {
		user = biz.GetUserFromSession(v.(string))
		if user.Id > 0 && strings.Index(user.Roles, "A") != -1 {
			c.Next()
		}else{
			if (req.Method == "POST") {
				r.JSON(200, map[string]interface{}{"code": 90000, "message" : "你不是管理员，无权进行此操作！"})
				return
			}
			r.Redirect("/")
			return
		}
	}else{
		if (req.Method == "POST") {
			r.JSON(200, map[string]interface{}{"code": 90000, "message" : "你不是管理员，无权进行此操作！"})
			return
		}
		r.Redirect("/")
	}
}