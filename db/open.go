package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conn() (*sql.DB, error) {
	// root:root@tcp(127.0.0.1:3306)/test 说明：
	// 第一个root是用户名
	// 第二个root是密码
	// 127.0.0.1:3306是地址
	// test 是数据库
	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	return conn, err
}
