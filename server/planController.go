package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	analyzer "grabContent/Cluster"
)

const (
	UserLovePara float64 = 0.75 
)

func fetchPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		sk := r.FormValue("key")
		// config := r.FormValue("config")

		if userMap[username].sk!= "" {
			if matchSessionKey(sk, userMap[username].sk){
				var content Content
				var infolist = make([]interface{}, 0, 10)
				var contentArr = make([]string, 0, 10)
				userInterests := stringToInterest(userMap[username].interest)
				for len(infolist)<=7 {
					tempContents := content.QueryRandom()
					for _, value := range tempContents {
						if !checkContentRepeat(value.contents["id"], contentArr) {
							if analyzer.GetInterestDegree(userInterests, stringToInterest(value.contents["tags"])) > UserLovePara {
								tempContent := map[string]interface{}{
								"content_id":value.contents["id"], 
								"title":value.contents["summary"],
								"summary":value.contents["summary"],
								"type": value.contents["type"],
								"cover_url":value.contents["cover_url"],
								"link":value.contents["link"],
								"author":value.contents["author"],
								"source":value.contents["source"],
								"tags":value.contents["tags"],
								"rates":value.contents["rates"],
								"like_num":value.contents["like_num"],
								"date" : value.contents["date"]}
								infolist = append(infolist, tempContent)
								contentArr = append(contentArr, value.contents["id"])
							} 
						}	
					}	
				}
				result := map[string]interface{}{"status": 0, "info_list": infolist}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}else {
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else {
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


func feedBack (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")
		// user



		result := map[string]string{"status": "0", "content": "是开发商副科级萨姆索诺夫马上飞，赛诺菲开始放假了书法家索科洛夫健康的快递师傅就开始地方就开始力"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

/*
	check the content does't repeat in the array
*/
func checkContentRepeat(contentId string, arr []string) bool{
	for _, value := range arr {
		if value == contentId {
			return true
		}
	}
	return false
}

/*
	transform interests to string
*/
func interestToString(interest map[string] float32) string {
	strResult,_ := json.Marshal(interest)
	return string(strResult)
}

/*
	transform string to interests
*/
func stringToInterest(str string) map[string] float32 {
	var dat map[string] float32
	if err := json.Unmarshal([]byte(str), &dat); err != nil {
        panic(err)
    }
    return dat
}

func fetchContentItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")

		result := map[string]string{"status": "0", "content": "是开发商副科级萨姆索诺夫马上飞，赛诺菲开始放假了书法家索科洛夫健康的快递师傅就开始地方就开始力"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

type Response1 struct {
    page   int
    fruits map[string]string
}

func fetchFavorList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		sk := r.FormValue("key")
		// start_offset := r.FormValue("start_offset")
		// number := r.FormValue("number")
		if userMap[username].sk!= "" {
			if matchSessionKey(sk, userMap[username].sk){
				var favor Favor
				var content Content
				// favorList := favor.QueryUserId(userMap[username].userId)
				favorList := favor.QueryUserId(1)
				var infolist = make([]interface{}, 0, 10)
				for _, value := range favorList {
					tempContent := content.QueryId(value.contents["content_id"])
					tempFavor := map[string]interface{}{
						"content_id":tempContent.contents["id"], 
						"title":tempContent.contents["summary"],
						"summary":tempContent.contents["summary"],
						"type": tempContent.contents["type"],
						"cover_url":tempContent.contents["cover_url"],
						"link":tempContent.contents["link"],
						"author":tempContent.contents["author"],
						"source":tempContent.contents["source"],
						"tags":tempContent.contents["tags"],
						"rates":tempContent.contents["rates"],
						"like_num":tempContent.contents["like_num"],
						"date" : value.contents["date"]}
					infolist = append(infolist, tempFavor)
				}
				result := map[string]interface{}{"status": 0, "info_list": infolist}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}else {
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else {
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

func toggleFavor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		sk := r.FormValue("key")
		contentId := r.FormValue("content_id")
		//0not favor  1 favored
		lastStatus := r.FormValue("last_status")
		if userMap[username].sk!= "" {
			if matchSessionKey(sk, userMap[username].sk){
				//do favor
				if lastStatus == "0" {
					var favor Favor
					favor.Init()
					favor.contents["user_id"] = userMap[username].userId
					// favor.contents["user_id"] = "1"
					favor.contents["content_id"] = contentId
					favor.contents["date"] = time.Now().Format("2006-01-02 15:04:05")
					if favor.insert() {
						result := map[string]string{"status": "0", "result": "成功"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}else{
						result := map[string]string{"status": "1", "result": "失败"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}
				}else {
					//cancel favor
					var favor Favor
					tempFavor := favor.QueryFavor(userMap[username].userId, contentId)
					// tempFavor := favor.QueryFavor("1", contentId)
					if tempFavor!=nil && tempFavor.delete() {
						result := map[string]string{"status": "0", "result": "成功"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}else {
						result := map[string]string{"status": "1", "result": "失败"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}
				}
			}else{
				//key not right
				result := map[string]string{"status": "2", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "3", "result": "请重新登录"}
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

func toggleLike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		sk := r.FormValue("key")
		contentId := r.FormValue("content_id")
		lastStatus := r.FormValue("last_status")
		
		if userMap[username].sk!= "" {
			if matchSessionKey(sk, userMap[username].sk){
				// do like
				if lastStatus == "0" {
					var likeContent LikeContent
					likeContent.Init()
					likeContent.contents["user_id"] = userMap[username].userId
					// likeContent.contents["user_id"] = "1"
					likeContent.contents["content_id"] = contentId
					likeContent.contents["date"] = time.Now().Format("2006-01-02 15:04:05")
					if likeContent.insert() {
						result := map[string]string{"status": "0", "result": "成功"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}else{
						result := map[string]string{"status": "1", "result": "失败"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}
				}else{
					//cancel like
					var likeContent LikeContent
					tempLikeContent := likeContent.QueryLikeContent(userMap[username].userId, contentId)
					// tempLikeContent := likeContent.QueryLikeContent("1", contentId)
					if tempLikeContent!=nil && tempLikeContent.delete() {
						result := map[string]string{"status": "0", "result": "成功"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}else{
						result := map[string]string{"status": "1", "result": "失败"}
						strResult,_ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}

				}

			}else{
				//key not right
				result := map[string]string{"status": "2", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "3", "result": "请重新登录"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func fetchLikeNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		sk := r.FormValue("key")
		contentId := r.FormValue("content_id")
		if userMap[username].sk!= "" {
			if matchSessionKey(sk, userMap[username].sk){
				var content Content
				tempContent:=content.QueryId(contentId)
				if tempContent!=nil {
					result := map[string]string{"status": "0", "number": tempContent.contents["like_num"]}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}else {
					result := map[string]string{"status": "1", "result": "失败"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}
			}else{
				//key not right
				result := map[string]string{"status": "2", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			// no login or server down
			result := map[string]string{"status": "3", "result": "请重新登录"}
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

func doILike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		sk := r.FormValue("key")
		contentId := r.FormValue("content_id")
		if userMap[username].sk!= "" {
			if matchSessionKey(sk, userMap[username].sk){
				var likeContent LikeContent
				tempContent := likeContent.QueryLikeContent(userMap[username].userId, contentId)
				if tempContent !=nil {
					//like
					result := map[string]string{"status": "0", "result": "1"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}else {
					//not like
					result := map[string]string{"status": "0", "result": "0"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}

			}else {
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
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