package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
	get all categories
*/
func getCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if len(interestCategoriesMap) > 0 {
			//categories exsit

		} else {
			var interestCategory InterestCategory
			tempCategories := interestCategory.QueryAll()
			for _, value := range tempCategories {
				addInterestCategories(value.contents["name"], value.contents["pic"])
			}

		}

		strCategoriesMap, err := json.Marshal(interestCategoriesMap)
		fmt.Println("str categories:" + string(strCategoriesMap))
		if err == nil {
			result := map[string]interface{}{"status": "0", "categories": string(strCategoriesMap)}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		} else {
			defer func() {
				result := map[string]string{"status": "1", "result": "失败"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}()
			panic(err.Error())
		}

	} else {

		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
	返回用户没有订阅的领域 //deprecated
*/
func getSubscribeCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				//cache all categoriesMap
				if len(interestCategoriesMap) > 0 {
					//categories exsit

				} else {
					var interestCategory InterestCategory
					tempCategories := interestCategory.QueryAll()
					for _, value := range tempCategories {
						addInterestCategories(value.contents["name"], value.contents["pic"])
					}
				}
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)

				var oldCategories map[string]int
				if err := json.Unmarshal([]byte(tempUser.contents["category"]), &oldCategories); err != nil {
					panic(err)
				}
				//add categories not in user
				newCategories := make(map[string]string)
				for cateKey, cateValue := range interestCategoriesMap {
					if _, ok := oldCategories[cateKey]; !ok {
						newCategories[cateKey] = cateValue
					}
				}
				strNewCategories, _ := json.Marshal(newCategories)
				result := map[string]interface{}{"status": "0", "categories": string(strNewCategories)}
				strResult, _ := json.Marshal(result)
				fmt.Println("subscribeCategories:", string(strResult))
				fmt.Fprintf(w, string(strResult))

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
	get user categories
*/
func getUserCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {

				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				result := map[string]string{"status": "0", "categories": tempUser.contents["category"]}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
set user categories
*/
func setUserCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		tempCategories := r.FormValue("categories")
		fmt.Println("set UserCategories" + tempCategories)
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				var oldCategories map[string]int
				if err := json.Unmarshal([]byte(tempUser.contents["category"]), &oldCategories); err != nil {
					panic(err)
				}
				for _, value := range oldCategories {
					if value == 0 {
						//degree zero is not allowed
						result := map[string]string{"status": "5", "result": "程度不能为零"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
						return
					}
				}
				tempUser.contents["category"] = tempCategories
				if tempUser.update() {
					result := map[string]string{"status": "0", "result": "成功"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]string{"status": "1", "result": "失败"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
add user categories  //deprecated
*/
func addUserCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		tempCategories := r.FormValue("categories")
		fmt.Println("addUserCategories" + tempCategories)
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				var oldCategories map[string]int
				var addCategories map[string]int
				if err := json.Unmarshal([]byte(tempUser.contents["category"]), &oldCategories); err != nil {
					panic(err)
				}
				if err := json.Unmarshal([]byte(tempCategories), &addCategories); err != nil {
					panic(err)
				}

				for cateKey, cateValue := range addCategories {
					if _, ok := oldCategories[cateKey]; !ok {
						oldCategories[cateKey] = cateValue
					} else {
						// subscrib categories had subscribed
						result := map[string]string{"status": "6", "result": "订阅了已经订阅的领域"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
						return
					}
				}
				categoryStr, _ := json.Marshal(oldCategories)
				tempUser.contents["category"] = string(categoryStr)
				if tempUser.update() {
					result := map[string]string{"status": "0", "result": "成功"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]string{"status": "1", "result": "失败"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
	get fchannle list for subscribe
*/
func getChannelList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		num := r.FormValue("number")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				//cache all the interest channels
				if len(interestChannelsMap) > 0 {
					//channels exsit
				} else {
					var interestList InterestList
					tempList := interestList.QueryAll()
					for _, value := range tempList {
						var interestCategory InterestCategory
						tempCategory := interestCategory.QueryInterestCategory(value.contents["category"])
						addInterestChannels(value.contents["name"], tempCategory.contents["name"],
							value.contents["pic"], value.contents["degree"])
					}

				}
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)

				var dat map[string]int
				var channelList = make([]interface{}, 0, 10)
				if err := json.Unmarshal([]byte(tempUser.contents["category"]), &dat); err != nil {
					panic(err)
				}
				if tempUser.contents["category"] == "{}" || tempUser.contents["category"] == "" || len(tempUser.contents["category"]) == 0 || len(dat) == 0 {
					for key1, value1 := range interestChannelsMap {
						tempChannel := map[string]interface{}{
							"name":   key1,
							"pic":    value1.pic,
							"degree": float64(value1.degree / float64(3))}
						channelList = append(channelList, tempChannel)
					}
				} else {
					//add all the user may interest channels
					for key, value := range dat {
						for key1, value1 := range interestChannelsMap {
							if key == value1.category {
								tempChannel := map[string]interface{}{
									"name":   key1,
									"pic":    value1.pic,
									"degree": value1.degree * float64(value+1) / 3}
								channelList = append(channelList, tempChannel)
							}
						}
					}
				}

				// fmt.Println("channelList:+")

				// for key,value := range channelList {
				// 	fmt.Println(key)
				// 	fmt.Println(value.(map[string]interface{})["degree"].(float64))
				// 	fmt.Println(channelList[key].(map[string]interface{})["degree"].(float64))
				// 	fmt.Println("-")
				// }

				//sort to find the max n
				tempLen := len(channelList)
				fmt.Println("channelList:+")
				fmt.Println(tempLen)
				for key, _ := range channelList {
					for j := key + 1; j < tempLen; j++ {
						// fmt.Print(value)
						// fmt.Print("::::")
						// fmt.Print(channelList[j].(map[string]interface{})["degree"].(float64))
						if channelList[key].(map[string]interface{})["degree"].(float64) <
							channelList[j].(map[string]interface{})["degree"].(float64) {
							temp := channelList[j]
							channelList[j] = channelList[key]
							channelList[key] = temp
						}
					}
					fmt.Println("-----------")
				}
				number := 10
				if num == "" {

				} else {
					number, _ = strconv.Atoi(num)
				}
				var resultList = channelList[:number]
				strList, _ := json.Marshal(resultList)
				fmt.Println(string(strList))
				result := map[string]interface{}{"status": "0", "channels": string(strList)}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
	get user interested channels
*/
func getUserChannels(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				result := map[string]string{"status": "0", "channels": userMap[username].interest}
				fmt.Println("get user channels:" + userMap[username].interest)
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
	modify user channels
*/
func setUserChannels(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		tempChannels := r.FormValue("channels")
		fmt.Println("user channels:" + tempChannels)
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				tempUser.contents["interest"] = tempChannels
				if tempUser.update() {
					//modify the cache
					addUser(username, userMap[username].userId, userMap[username].sk, tempChannels, userMap[username].date)

					result := map[string]string{"status": "0", "result": "成功"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]string{"status": "1", "result": "失败"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
	add user channels
*/
func addUserChannels(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		tempChannels := r.FormValue("channels")
		fmt.Println("user channels:" + tempChannels)
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				if tempUser.contents["interest"] == "" || tempUser.contents["interest"] == "NULL" || tempUser.contents["interest"] == "{}" ||
					tempUser.contents["interest"] == "null" || len(tempUser.contents["interest"]) == 0 {
					tempUser.contents["interest"] = tempChannels
				} else {
					tempInterests := stringToInterest(tempUser.contents["interest"])
					outInterests := stringToInterest(tempUser.contents["interest"])
					tempAddInterests := stringToInterest(tempChannels)
					for key, _ := range tempInterests {
						for key1, value1 := range tempAddInterests {
							if key != key1 {
								outInterests[key1] = value1
							}
						}
					}
					tempUser.contents["interest"] = interestToString(outInterests)
					fmt.Println("added channels:" + interestToString(outInterests))
				}
				if tempUser.update() {
					//modify the cache
					addUser(username, userMap[username].userId, userMap[username].sk, tempUser.contents["interest"], userMap[username].date)

					result := map[string]string{"status": "0", "result": "成功"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]string{"status": "1", "result": "失败"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}

			} else {
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}
