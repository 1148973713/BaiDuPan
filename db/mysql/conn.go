package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(119.23.57.189:3036)/fileserver?charset=utf8")
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

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	//返回列名
	columns, _ := rows.Columns()
	//https://www.cnblogs.com/echojson/p/10746807.html
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	for j := range values {
		//&符号的意思是对变量取地址，如：变量a的地址是&a
		//*符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		//rows.Scan自动赋值
		err := rows.Scan(scanArgs...)
		checkErr(err)

		//便利values数组，如果不为空，则将index传入columns,record[columns[i]]实现了列名与数值相对应：password - 123456  7891011
		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		fmt.Println(record)
		//形成数组
		records = append(records, record)
	}

	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
