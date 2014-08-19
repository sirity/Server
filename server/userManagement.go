package server

import (

)

type UserPie struct {
	sk string
	interest string
	date string
}

var userMap map[string]UserPie
var userHung map[string]string

func ManageInit(){
	userMap = make(map[string]UserPie)
}

func addUser(username string, sk string, interest string, date string) {
	userMap[username] = UserPie{sk:sk, interest:interest, date:date}
}

func logOut(username string) {
	delete(userMap, username)
}
