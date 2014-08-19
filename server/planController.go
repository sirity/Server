package server

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func fetchPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")

		tag := []string{"math", "english", "football"}
		var infolist [2] interface{}
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

func fetchFavorList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		// username := r.FormValue("username")
		// sk := r.FormValue("key")
		// contentId := r.FormValue("content_id")
		tag := []string{"math", "english", "football"}
		var infolist [2] interface{}
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

func toggleFavor(w http.ResponseWriter, r *http.Request) {
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