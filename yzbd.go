package main

import (
	"time"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"route"
	"html/template"
	"services"
	//"yztest"
)

func main()  {
	// now,err := time.Parse("2006-01-02 15:04:05 +0000 UTC", "2018-01-09 11:02:51 +0000 UTC" )
	// if err != nil {
	// 	common.Debug(err)
	// }
	// common.Debug(now.Format("2006-01-02 15:04:05"))
	services.Debug("main started")
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:"templates",
		Layout: "layout",
		Extensions:[]string{".tmpl",".html"},
		Funcs: []template.FuncMap{
			{
				"unescaped": func(x string) template.HTML {
					return template.HTML(x)
				},
				"formatToDate" : func(x time.Time) template.HTML {
					return template.HTML(x.Format("2006-01-02"))
				},
				"formatToTime" : func(x time.Time) template.HTML {
					return template.HTML(x.Format("2006-01-02 15:04:05"))
				},
			},
		},
		Charset:"UTF-8",
		IndentJSON: true,
	}))
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	route.Run(m)
}
