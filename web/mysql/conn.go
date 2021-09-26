package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//初始化连接
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(mysql:3306)/admin")
	if err != nil {
		panic(err)
	}

}
