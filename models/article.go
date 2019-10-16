package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`   //标题
	Desc       string `json:"desc"`    //简介
	Content    string `json:"content"` //内容
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//按id获取文章
func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

//取得文章总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

//获取文章

/*Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID的方式去找到这两个类之间的关联关系
Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询*/
func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	//Related  相关
	db.Model(&article).Related(&article.Tag) //查找与文章相关的文章标签
	return
}

//获取全部文章
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

//修改文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Update(data)
	return true

}

//删除文章
func DeleteArticle(id int) bool {
	db.Where("id=?", id).Delete(Article{})
	return true
}

//添加文章
//INSERT  INTO `blog_article` (`created_on`,`modified_on`,`tag_id`,`title`,`desc`,`content`,`created_by`,`modified_by`,`state`) VALUES (1571125572,0,2,'test1','test-desc','test-content','test-created','',1)
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on!=?", 0).Delete(&Article{})
	return true
}
