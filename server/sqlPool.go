package server

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	userTable string = "user"
)

func Init(){
	db1, err := sql.Open("mysql", "root:123@/mercury")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    db = db1
}
