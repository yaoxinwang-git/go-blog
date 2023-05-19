package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	// 获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")

	// 7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)

	if err != nil {
		detail.WriteError(w, errors.New("不识别此路径请求"))
		return
	}

	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, err)
		return
	}
	detail.WriteData(w, postRes)
}
