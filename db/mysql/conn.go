package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/fileserver?charset=utf8mb4")
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(2 << 10)
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

// DBConn 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
