package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("msql", "root:root@tcp(119.23.57.189:3036)/fileserver?charset=utf8")
	db.SetMaxIdleConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Print("Filed to connect to mysql,err:" + err.Error())
		os.Exit(1)
	}
}

//DBCoonn：返回数据库连接对象
func DBCoon() *sql.DB {
	return db
}
