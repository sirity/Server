package server

import (

)

type UserPie struct {
	userId string
	sk string
	interest string
	date string
}

// func (userPie *UserPie) setInterest(interest string) {
// 	userPie.interest = interest
// }

var userMap map[string]UserPie
var userHung map[string]string
var userForget map[string]string

func ManageInit(){
	userMap = make(map[string]UserPie)
	userHung = make(map[string]string)
	userForget = make(map[string]string)
}

func addUser(username string, userId string, sk string, interest string, date string) {
	userMap[username] = UserPie{userId:userId, sk:sk, interest:interest, date:date}
}



// func modifyUserInterest(username, interest string) {
// 	userMap[username].interest = interest
// }

func logOut(username string) {
	delete(userMap, username)
}

func addHung(username, ak string) {
	userHung[username] = ak
}

func deleteHung(username string) {
	delete(userHung, username)
}

func addForget(username, fk string) {
	userForget[username] = fk
}

func deleteForget(username string) {
	delete(userForget, username)
}
