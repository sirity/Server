package webserver

/*
	this package is a webproxy of sirity server
	aim to show result on webbrowser
	Realization way is proxy for package server
*/
import (
	"fmt"
	"log"
	"net/http"
)

var mymux *http.ServeMux

func Run() {

	mymux = http.NewServeMux()
	//绑定路由
	bind()
	err := http.ListenAndServe(":3434", mymux) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func bind() {
	mymux.HandleFunc("/", login)
}
