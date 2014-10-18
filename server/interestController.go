package server

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func getCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		if len(interestCategoriesMap) > 0{
			//categories exsit
			
		}else{
			var interestCategory InterestCategory
			tempCategories := interestCategory.QueryAll()
			for _, value := range tempCategories {
				addInterestCategories(value.contents["name"], value.contents["pic"])
			}
			
		}
		strCategoriesMap,err := json.Marshal(interestCategoriesMap)
		if err ==nil {
			result := map[string]string{"status": "0", "result": string(strCategoriesMap)}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}else{
			defer func() {
				result := map[string]string{"status": "1", "result": "失败"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}()
			panic(err.Error())
		}
		
	}else{
		//network wrong
		result := map[string]string{"status": "2", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func getUserCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk!= "" {
			if matchSessionKey(key, userMap[username].sk){

				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
    			result := map[string]string{"status": "0", "categories": tempUser.contents["category"]}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))

			}else{
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	}else{
		//network wrong
		result := map[string]string{"status": "3", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func setUserCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		key := r.FormValue("key")
		tempCategories := r.FormValue("categories")
		if userMap[username].sk!= "" {
			if matchSessionKey(key, userMap[username].sk){

				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				tempUser.contents["category"] = tempCategories
				if tempUser.update() {
	    			result := map[string]string{"status": "0", "result": "成功"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
	    			result := map[string]string{"status": "0", "result": "失败"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}

			}else{
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	}else{
		//network wrong
		result := map[string]string{"status": "3", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func getChannelList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk!= "" {
			if matchSessionKey(key, userMap[username].sk){
				//cache all the interest channels
				if len(interestChannelsMap) > 0{
					//channels exsit
				}else{
					var interestList InterestList
					tempList := interestList.QueryAll()
					for _, value := range tempList {
						addInterestChannels(value.contents["name"], value.contents["category"],
						 value.contents["pic"], value.contents["degree"])
					}
					
				}
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				var dat map[string] int
				if err := json.Unmarshal([]byte(tempUser.contents["category"]), &dat); err != nil {
        			panic(err)
    			}

    			//add all the user may interest channels
    			var channelList = make([]interface{}, 0, 10)
    			for key, value := range dat {
    				for key1, value1 := range interestChannelsMap{
    					if key == value1.category {
    						tempChannel := map[string]interface{}{
    						"name": key1,
    						"pic": value1.pic,
    						"degree": value1.degree * float64(value)}
    						channelList = append(channelList, tempChannel)
    					}
    				}
    			}

    			//sort to find the max n
    			tempLen := len(channelList)
    			for key, value := range channelList {
    				for j:= key+1; j<tempLen; j++ {
    					if value.(map[string]interface{})["degree"].(float64) < 
    						channelList[j].(map[string]interface{})["degree"].(float64) {
    						temp := channelList[j]
    						channelList[j] = value
    						value = temp
    					}
    				}
    			}

    			var resultList = channelList[:9]
    			result := map[string]interface{}{"status": "0", "channel_list": resultList}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))

			}else{
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	}else{
		//network wrong
		result := map[string]string{"status": "3", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func getUserChannels(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk!= "" {
			if matchSessionKey(key, userMap[username].sk){
    			result := map[string]string{"status": "0", "channel_list": userMap[username].interest}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))

			}else{
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	}else{
		//network wrong
		result := map[string]string{"status": "3", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func setUserChannels(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		key := r.FormValue("key")
		tempChannels := r.FormValue("channels")
		if userMap[username].sk!= "" {
			if matchSessionKey(key, userMap[username].sk){
				//get user categories
				var user User
				tempUser := user.QueryId(userMap[username].userId)
				tempUser.contents["interest"] = tempChannels
				if tempUser.update() {
					//modify the cache
					addUser(username, userMap[username].userId, userMap[username].sk, tempChannels, userMap[username].date)

	    			result := map[string]string{"status": "0", "result": "成功"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}else {
	    			result := map[string]string{"status": "1", "result": "失败"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}

			}else{
				//key not right
				result := map[string]string{"status": "3", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "2", "result": "请重新登录"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	}else{
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

