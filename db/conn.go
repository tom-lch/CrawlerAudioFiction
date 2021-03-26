package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConn(Name, Pwd, Host, Port, DBname string) *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Name, Pwd, Host, Port, DBname)
	DB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("connect mysql failed, checkout dataSourceName: ", err)
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Fatal("open database fail")
		return nil
	}
	return DB
}
