package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Post struct {
	Id        int
	Title     string
	Content   string
	Author    string `sql:"not null"`
	CreatedAt time.Time
	Comments  []Comment
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var DbConn *gorm.DB

func init() {
	var err error
	DbConn, err = gorm.Open("mysql", "root:123456@tcp(mysql:3306)/admin")
	if err != nil {
		panic(err)
	}
	DbConn.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Title: "GORM 示例教程", Content: "基于 GORM 进行数据库增删改查", Author: "学院君"}
	// 通过 GORM 插入文章记录
	DbConn.Create(&post)
	fmt.Println(post)

	// 通过关联关系新增评论并将其附加到对应的文章记录
	comment := Comment{Content: "Test Comment", Author: "学院君小号"}
	DbConn.Model(&post).Association("Comments").Append(comment)

	// 查询文章记录
	var gormPost Post
	fmt.Println(gormPost)
	DbConn.Where("author = ?", "学院君").First(&gormPost)
	fmt.Println(gormPost)

	// 查询包含评论数据的文章记录
	var comments []Comment
	DbConn.Model(&gormPost).Related(&comments)
	fmt.Println(comments)

}
