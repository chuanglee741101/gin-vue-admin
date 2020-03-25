package service

import (
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//IArticleService Article接口定义
type IArticleService interface {
	//GetArticle 根据id获取Article
	GetArticle(id int) *models.Article
	//GetTables 分页返回文章
	GetTables(page, pagesize int) *[]models.Article
	//AddArticle 新增Article
	AddArticle(article *models.Article) bool
	// //ExistArticleByID 根据ID判断Article是否存在
	// ExistArticleByID(id int) bool
	// //EditArticle 编辑Article
	// EditArticle(article models.Article) bool
	// //DeleteArticle 删除Article
	// DeleteArticle(id int) bool
	//GetArticles 获取文章信息
	GetArticles(PageNum int, PageSize int, total *uint64, maps interface{}) *[]models.Article
}
