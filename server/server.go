package server

import (
	"net/http"
	"log"
    "fmt"
)

var mymux *http.ServeMux
const (
    mailAddress string = "http://121.40.190.238:1280"
)

func Run() {
	
	mymux = http.NewServeMux()
    //绑定路由
    bind()

    err := http.ListenAndServe(":1280", mymux) //设置监听的端口

    if err != nil {

        log.Fatal("ListenAndServe: ", err)
    }
}

func TestDatabase() {
    var user User
    users := user.QueryAll()
    var le = len(users)
    for i := 0; i < le; i++ {
        fmt.Println(users[i].contents["id"])
        fmt.Println(users[i].contents["username"])
        fmt.Println("xxxxxxxxxxxxxxx")
    }
}

func TestUsers() {
    var user User
    users := user.QueryAll()
    var le = len(users)
    for i := 0; i < le; i++ {
        fmt.Println(users[i].contents["id"])
        fmt.Println(users[i].contents["username"])
        fmt.Println("xxxxxxxxxxxxxxx")
    }
}