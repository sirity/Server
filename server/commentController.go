package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func fetchCommentList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		sk := r.FormValue("key")
		contentId := r.FormValue("content_id")
		// startOffset := r.FormValue("start_offset")
		// number := r.FormValue("number")

		if userMap[username].sk != "" {
			if matchSessionKey(sk, userMap[username].sk) {
				var comment Comment
				commentMap := comment.QueryContentComment(contentId)
				var infolist = make([]interface{}, 0, 10)
				for _, value := range commentMap {
					tempComment := map[string]interface{}{
						"comment_id":    value.contents["id"],
						"body":          value.contents["body"],
						"username":      value.contents["user_id"],
						"to_comment_id": value.contents["to_comment_id"],
						"date":          value.contents["date"],
						"like_num":      value.contents["like_num"]}
					infolist = append(infolist, tempComment)
				}
				result := map[string]interface{}{"status": 0, "comment_list": infolist}
				strResult, _ := json.Marshal(result)
				fmt.Println("comment_list:" + string(strResult))
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

func postComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		contentId := r.FormValue("content_id")
		commentStr := r.FormValue("comment")
		fmt.Println("commentId" + contentId)
		fmt.Println("comment" + commentStr)

		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				var comment Comment
				comment.Init()
				var dat map[string]string
				err := json.Unmarshal([]byte(commentStr), &dat)
				if err == nil {
					comment.contents["content_id"] = contentId
					comment.contents["user_id"] = userMap[username].userId
					comment.contents["to_comment_id"] = dat["to_comment_id"]
					comment.contents["body"] = dat["body"]
					comment.contents["date"] = dat["date"]
					comment.contents["like_num"] = "0"
					if comment.insert() {
						result := map[string]interface{}{"status": 0, "result": "成功"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					} else {
						result := map[string]interface{}{"status": 1, "result": "失败"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}
				} else {
					panic(err)
					defer func() {
						result := map[string]interface{}{"status": 5, "result": "数据格式错误"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}()
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

func deleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		commentId := r.FormValue("comment_id")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				var comment Comment
				tempComment := comment.QueryId(commentId)
				if tempComment.delete() {
					result := map[string]interface{}{"status": "0", "result": "删除成功"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]interface{}{"status": "1", "result": "删除失败"}
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

//TODO-comment like add or cut
func commentToggleLike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		sk := r.FormValue("key")
		commentId := r.FormValue("comment_id")
		lastStatus := r.FormValue("last_status")
		if userMap[username].sk != "" {
			if matchSessionKey(sk, userMap[username].sk) {
				// do like
				if lastStatus == "0" {
					var likeComment LikeComment
					likeComment.Init()
					likeComment.contents["user_id"] = userMap[username].userId
					likeComment.contents["comment_id"] = commentId
					likeComment.contents["date"] = time.Now().Format("2006-01-02 15:04:05")
					if likeComment.insert() {
						result := map[string]interface{}{"status": "0", "result": "成功"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					} else {
						result := map[string]interface{}{"status": "1", "result": "失败"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}
				} else {
					//cancel like
					var likeComment LikeComment
					tempLikeComment := likeComment.QueryLikeComment(userMap[username].userId, commentId)
					if tempLikeComment.delete() {
						result := map[string]interface{}{"status": "0", "result": "成功"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					} else {
						result := map[string]interface{}{"status": "1", "result": "失败"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}

				}
				var comment Comment
				tempComment := comment.QueryId(commentId)
				if tempComment.delete() {
					result := map[string]interface{}{"status": "0", "result": "删除成功"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]interface{}{"status": "1", "result": "删除失败"}
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

func commentFetchLikeNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		sk := r.FormValue("key")
		commentId := r.FormValue("comment_id")
		if userMap[username].sk != "" {
			if matchSessionKey(sk, userMap[username].sk) {
				var comment Comment
				tempComment := comment.QueryId(commentId)
				if tempComment != nil {
					result := map[string]interface{}{"status": "0", "number": tempComment.contents["like_num"]}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]interface{}{"status": "1", "result": "失败"}
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

func commentDoILike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		sk := r.FormValue("key")
		commentId := r.FormValue("comment_id")
		if userMap[username].sk != "" {
			if matchSessionKey(sk, userMap[username].sk) {
				var likeComment LikeComment
				tempLikeComment := likeComment.QueryLikeComment(userMap[username].userId, commentId)
				if tempLikeComment != nil {
					result := map[string]interface{}{"status": "0", "result": "1"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]interface{}{"status": "1", "result": "0"}
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
		result := map[string]string{"status": "0", "result": "成功"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	} else {
		//network wrong
		result := map[string]string{"status": "4", "result": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}
