package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)

	if err != nil {

		categoryTemplate.WriteError(w, errors.New("不识别此路径请求"))
		return

	}
	// 处理分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}

	pageStr := r.Form.Get("page")

	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10

	//数据库查询
	categoryResponse,err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w,err)
		return 
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
