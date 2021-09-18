package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      int
	Title   string
	Content string
	Author  string
}

//获取单条数据
func GetPost(id int) (post Post, err error) {
	rows, err := Db.Query("select id, title, content, author from posts where id = ? limit 1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close() //最后关闭连接
	for rows.Next() {
		//post = Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author)
		if err != nil {
			panic(err)
		}
	}
	return
}

var Db *sql.DB

func init() {
	var err error
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//username:password@protocol(address)/dbname?param=value
	Db, err = sql.Open("mysql", "root:123456@tcp(mysql:3306)/admin")
	//Db, err = sql.Open("mysql", "b2b_dev:B2b.dev2021@tcp(db.b2b.orsd.tech:3306)/yzc_dev")
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 2, Title: "Go 语言数据库操作", Content: "基于第三方 go-sql-driver/mysql 包实现 MySQL 数据库增删改查", Author: "学院君"}

	// 获取单条记录
	dbPost, _ := GetPost(post.Id)
	fmt.Println(dbPost)

}
