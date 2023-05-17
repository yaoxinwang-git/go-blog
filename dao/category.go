package dao

import (
	"log"
	"go-blog/models"
)

func GetCategoryNameById(categoryId int)string{
	row:=DB.QueryRow("select name from blog_category where cid = ?",categoryId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category,error){
	rows ,err := DB.Query("select * from blog_category")
	if err !=  nil {
		log.Println("GetAllCategory 查询出错:",err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid,&category.Name,&category.CreateAt,&category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:",err)
			return nil, err
		}
		categorys = append(categorys,category)
	}
	return categorys,nil   
}
