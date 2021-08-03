package db

import (
	"BaiDuPan/db/mysql"
	"fmt"
)

//注册操作
func UserSignup(username string, password string) bool {
	stmt, err := mysql.DBCoon().Prepare("insert ignore into tbl_user(`user_name`,`user_pwd`)value (?,?)")
	if err != nil {
		fmt.Print("Failed to instert,err" + err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("Failed to instert,err" + err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}
	return false
}

func UserSignIn(username string, encpwd string) bool {
	stmt, err := mysql.DBCoon().Prepare("select * from tbl_user where user_name=?limit1")
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}
	row, err := stmt.Query(username)
	if err != nil {
		fmt.Printf(err.Error())
		return false
	} else if row == nil {
		fmt.Printf("user not find:" + username)
		return false
	}
	defer stmt.Close()
	pProws := mysql.ParseRows(row)
}
