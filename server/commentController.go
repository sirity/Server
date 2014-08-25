package server

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func fetchCommentList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")
		// startOffset := r.FormValue("start_offset")
		// number := r.FormValue("number")

		var infolist [4] interface{}
		infolist[0] = map[string]interface{}{
			"comment_id":"12340293",
			"body":"这篇文章不错哦",
			"username":"626731494@qq.com",
			"to_comment_id":"1232131",
			"date":"2012-3-23",
			"like_num":"12"}
		infolist[1] = map[string]interface{}{
			"comment_id":"12340293",
			"body":"这篇文章不错哦",
			"username":"626731494@qq.com",
			"to_comment_id":"1232131",
			"date":"2012-3-23",
			"like_num":"12"}
		infolist[2] = map[string]interface{}{
			"comment_id":"12340293",
			"body":"这篇文章不错哦",
			"username":"626731494@qq.com",
			"to_comment_id":"1232131",
			"date":"2012-3-23",
			"like_num":"12"}
		infolist[3] = map[string]interface{}{
			"comment_id":"12340293",
			"body":"这篇文章不错哦",
			"username":"626731494@qq.com",
			"to_comment_id":"1232131",
			"date":"2012-3-23",
			"like_num":"12"}

		result := map[string]interface{} {"status": "0", "comment_list": infolist}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func postComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// key := r.FormValue("key")
		// comment := r.FormValue("comment")

		result := map[string]string{"status": "0", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// key := r.FormValue("key")
		// contentId := r.FormValue("content_id")

		result := map[string]interface{}{"status": "0", "result":"删除成功"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}