package main

import (
	"database/sql"
	"fmt"

	//mysqlドライバ go get -u github.com/go-sql-driver/mysql でインストールする
	_ "github.com/go-sql-driver/mysql"
)

//DBを用意
var Db *sql.DB

//DB設定
func init() {
	var err error
	//DB構造体に設定を記述(接続しているわけではない)
	Db, err = sql.Open("mysql", "root:ohs80340@tcp(127.0.0.1:3306)/world")
	if err != nil {
		fmt.Println("DB INIT ERROR:", err.Error())
		return
	}
}
