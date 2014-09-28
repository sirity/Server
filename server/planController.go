package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)

func fetchPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		sk := r.FormValue("key")
		config := r.FormValue("config")

		fmt.Println(username)
		fmt.Println(sk)
		fmt.Println(config)

		tag := []string{"math", "english", "football"}
		var infolist [3] interface{}
		infolist[0] = map[string]interface{}{
			"content_id":"1000001", 
			"title":"ABC是什么意思",
			"summary":"ABC到底是什么意思",
			"type":"int",
			"cover_url":"http://static.cnbetacdn.com/newsimg/2014/0819/29_1jYFp2Rql.jpg",
			"link":"http://www.cnbeta.com/articles/320323.htm",
			"author":"danis",
			"source":"huxiu",
			"tags":tag,
			"rates":"3.5",
			"like_num":"12",
			"do_i_like":"1",
			"do_i_favor":"1"}
		infolist[1] = map[string]interface{}{
			"content_id":"1000001",
			"title":"ABC是什么意思",
			"summary":"ABC到底是什么意思",
			"type":"int",
			"cover_url":"http://static.cnbetacdn.com/newsimg/2014/0819/29_1jYFo0kK9.jpg",
			"link":"http://www.cnbeta.com/articles/320321.htm",
			"author":"danis",
			"source":"huxiu",
			"tags":tag,
			"rates":"3.5",
			"like_num":"12",
			"do_i_like":"1",
			"do_i_favor":"1"}
		infolist[2] = map[string]interface{}{
			"content_id":"1000001",
			"title":"ABC是什么意思",
			"summary":"ABC到底是什么意思",
			"type":"int",
			"cover_url":"http://www.ituring.com.cn/download/01fwlV6fUtOu",
			"link":"http://www.ituring.com.cn/article/119957",
			"author":"danis",
			"source":"huxiu",
			"tags":tag,
			"rates":"3.5",
			"like_num":"12",
			"do_i_like":"1",
			"do_i_favor":"1"}
		result := map[string]interface{}{"status": "0", "info_list": infolist}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
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
					number64, err := strconv.ParseInt(value.contents["content_id"], 10, 32)
		            var number int = int(number64)
		            checkErr(err)
					tempContent := content.QueryId(number)
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
		// contentId := r.FormValue("content_id")

		if userMap[username].sk!= "" {
			if matchSessionKey(sk, userMap[username].sk){

				result := map[string]string{"status": "0", "result": "成功"}
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

func toggleLike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")

		result := map[string]string{"status": "0", "result": "成功"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func fetchLikeNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")

		result := map[string]string{"status": "0", "number": "233"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func doILike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")
		
		result := map[string]string{"status": "0", "result": "成功"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}