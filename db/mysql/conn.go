package mysql

import (
	"database/sql"
	"fmt"
	"myCloud/config"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init(){
	db,_=sql.Open("mysql", config.MySQLSource )
	db.SetMaxOpenConns(10)
	err:=db.Ping()
	if err != nil{
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}


// DBConn : 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}