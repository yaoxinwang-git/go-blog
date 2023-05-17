package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router()  {
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	http.HandleFunc("/",views.HTML.Index)
	
	http.HandleFunc("/c/",views.HTML.Category)
	http.HandleFunc("/login",views.HTML.Login)
	
	http.HandleFunc("/api/v1/post",api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/login",api.API.Login)
	http.Handle("/resource/",http.StripPrefix("/resource/",http.FileServer(http.Dir("public/resource/"))))
}
