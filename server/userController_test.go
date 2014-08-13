package server

import (
	"testing"
	"fmt"
    "net/http"
    "net/url" 
    "io/ioutil"
)

func TestLogin(t *testing.T){
	Run()
	Init()
	ManageInit()
	resp, err := http.PostForm("http://127.0.0.1:1280/user/login",
        url.Values{"username": {"test"}, "password": {"123"}, "date": {"2014-8-10"}, "random": {"1299"}})
    if err != nil {
        fmt.Println(err)
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 
    fmt.Println(string(body))
}

func BenchmarkLogin(b *testing.B) {
	b.StopTimer()
	Init()
	b.StartTimer()
	var user User
	for i := 0; i < b.N; i++ {
		user.QueryAll()
	}

}
