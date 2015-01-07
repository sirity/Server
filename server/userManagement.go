package server

/*
	this file is design to maintain the cache of the whole system
	include the user sessions, interests, channels
*/
import (
	"strconv"
)

type UserPie struct {
	userId string
	sk string
	interest string
	date string
}

type InterestChannel struct {
	category string
	pic string
	degree float64
}

// func (userPie *UserPie) setInterest(interest string) {
// 	userPie.interest = interest
// }

//user session map
var userMap map[string]UserPie
//user active session
var userHung map[string]string
//user forget and find session
var userForget map[string]string
//interest categories cache
var interestCategoriesMap map[string]string
//interest channels cache
var interestChannelsMap map[string] InterestChannel

func ManageInit(){
	userMap = make(map[string]UserPie)
	userHung = make(map[string]string)
	userForget = make(map[string]string)
	interestCategoriesMap = make(map[string]string)
	interestChannelsMap = make(map[string]InterestChannel)
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

func addInterestCategories(name, link string) {
	interestCategoriesMap[name] = link
}

func addInterestChannels(name, category, pic, degree string) {
	temp, err := strconv.ParseFloat(degree, 64)
	if err == nil {
		interestChannelsMap[name] = InterestChannel{category: category, pic: pic, degree: temp}
	}else{
		panic(err.Error())
	}
}


