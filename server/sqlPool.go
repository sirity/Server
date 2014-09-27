package server

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var db *sql.DB

const (
	userTable string = "user"
	contentTable string = "content"
	commentTable string = "comment"
	favorTable string = "favor"
	likeCommentTable string = "likeComment"
	likeContentTable string = "likeContent"
)

func Init(){
	db1, err := sql.Open("mysql", "root:33a55c67@tcp(127.0.0.1:3306)/mercury")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    db = db1
}

func checkErr(err error) {
    if err != nil {
    	log.Println(err)
        panic(err)
    }
}
