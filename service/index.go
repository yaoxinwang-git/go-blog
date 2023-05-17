package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, _ := dao.GetPostPage(page, pageSize)
	//处理数据 post->postMore
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)

		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid: post.Pid,
			Title: post.Title,
			Slug: post.Slug,
			Content: template.HTML(content),
			CategoryId: post.CategoryId,
			CategoryName: categoryName,
			UserId: post.UserId,
			UserName: userName,
			ViewCount: post.ViewCount,
			Type: post.Type,
			CreateAt: models.DateDay(post.CreateAt),
			UpdateAt: models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	//处理分页
	total := dao.CountGetAllPost()
	pagesCount := (total-1)/10+1
	var pages []int
	for i :=0 ;i< pagesCount; i++	{
		pages = append(pages,i+1)
	}

	var hr = &models.HomeResponse{
		Viewer: config.Cfg.Viewer,
		Categorys: categorys,
		Posts: postMores,
		Total: total,
		Page: page,
		Pages: pages,
		PageEnd: page!=pagesCount,
	}
	
	return hr, nil
}
