package db

import (
	"fmt"
	"myCloud/db/mysql"
)

func UserSignUp(username string, passwd string)bool {
	stmt,err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user('user_name','user_pwd')values(?,?)")
	if err != nil{
		fmt.Println("Failed to insert,err"+err.Error())
		return false
	}
	defer stmt.Close()

	ret,err := stmt.Exec(username, passwd)
	if err!= nil{
		fmt.Println("Failed to insert, err:"+err.Error())
		return false
	}
	if rowsAffected,err := ret.RowsAffected();nil==err && rowsAffected > 0{
		return true
	}
	return false
}
