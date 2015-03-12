package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

const (
	userTable             string = "user"
	contentTable          string = "content"
	commentTable          string = "comment"
	favorTable            string = "favor"
	likeCommentTable      string = "likeComment"
	likeContentTable      string = "likeContent"
	interestCategoryTable string = "interestCategory"
	interestListTable     string = "interestList"
)

func Init() {
	db1, err := sql.Open("mysql", "root:33a55c67@tcp(127.0.0.1:3306)/mercury")
	if err != nil {
		checkErr(err) // Just for example purpose. You should use proper error handling instead of panic
	}
	if db1 == nil {
		fmt.Println("数据库链接不到")
	}
	db = db1
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		// panic(err)
	}
}
