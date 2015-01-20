package server

import (
	"Server/image"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		date := r.FormValue("date")
		random := r.FormValue("random")
		if !CheckAttack(username, date, random) {
			var user User
			user1 := user.QueryUser(username)
			if user1.contents != nil {
				//user exist
				temp := []byte(password)
				h := md5.New()
				h.Write([]byte(user1.contents["password"] + date))
				if HexEncodeEqual(temp, h.Sum(nil)) {
					//login success
					sk := SessionKey(username)
					addUser(username, user1.contents["id"], sk, user1.contents["interest"], date)
					result := map[string]string{"status": "0", "key": sk}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					//password fail
					result := map[string]string{"status": "1", "key": "密码错误"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}
			} else {
				//user not exist
				result := map[string]string{"status": "8", "key": "用户名不存在"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			//it's not our client
			result := map[string]string{"status": "6", "key": "这不是我们的"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}
	} else {
		//network wrong
		result := map[string]string{"status": "4", "key": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}

}

// func register(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST"{
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")
// 		date := r.FormValue("date")
// 		random := r.FormValue("random")
// 		if !CheckAttack(username, date, random) {
// 			var user User
// 			user1 := user.QueryUser(username)
// 			if user1.contents !=nil{
// 				//user exist
// 				result := map[string]string{"status": "1", "result": "用户名存在"}
// 				strResult,_ := json.Marshal(result)
// 				fmt.Fprintf(w, string(strResult))
// 			}else{
// 				//user not exist
// 				var user User
// 				user.Init()
// 				user.contents["username"] = username
// 				user.contents["nickname"] = "我是小白"
// 				user.contents["password"] = password
// 				user.contents["status"] = "0"
// 				if(user.insert()){
// 					//insert success
// 					sk := SessionKey(username)
// 					//TO-DO bugs
// 					addUser(username, user.contents["id"], sk, user1.contents["interest"], date)

// 					ak := activeKey(username, date, random)
// 					addHung(username, ak)

// 					// SendRegisterMail(username, "http://127.0.0.1:1280/user/active/?username=" + username +
// 					// 	"&activekey=" + ak)
// 					SendRegisterMail(username, mailAddress + "/user/active/?username=" + username +
// 						"&activekey=" + ak)

// 					result := map[string]string{"status": "0", "key": sk}
// 					strResult,_ := json.Marshal(result)
// 					fmt.Fprintf(w, string(strResult))
// 				}else{
// 					//insert fail
// 					result := map[string]string{"status": "5", "key": "未知原因"}
// 					strResult,_ := json.Marshal(result)
// 					fmt.Fprintf(w, string(strResult))
// 				}
// 			}
// 		}else{
// 			//it's not our client
// 			result := map[string]string{"status": "4", "key": "这不是我们的"}
// 			strResult,_ := json.Marshal(result)
// 			fmt.Fprintf(w, string(strResult))
// 		}
// 	}else{
// 		//network wrong
// 		result := map[string]string{"status": "3", "key": "网络嗝屁了"}
// 		strResult,_ := json.Marshal(result)
// 		fmt.Fprintf(w, string(strResult))
// 	}
// }

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		date := r.FormValue("date")
		random := r.FormValue("random")
		fmt.Println("register")
		if !CheckAttack(username, date, random) {
			var user User
			user1 := user.QueryUser(username)
			if user1.contents != nil {
				//user exist
				result := map[string]string{"status": "1", "result": "用户名存在"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			} else {
				//user not exist
				var user User
				user.Init()
				user.contents["username"] = username
				user.contents["nickname"] = "我是小白"
				user.contents["password"] = password
				user.contents["status"] = "0"
				user.contents["category"] = "{}"
				user.contents["interest"] = "{}"
				if user.insert() {
					//insert success
					sk := SessionKey(username)
					//TO-DO bugs
					addUser(username, user.contents["id"], sk, user.contents["interest"], date)

					// ak := activeKey(username, date, random)
					// addHung(username, ak)
					tempCheckCode := getSpecificLenRandom(4)
					addHung(username, tempCheckCode)

					// SendRegisterMail(username, "http://127.0.0.1:1280/user/active/?username=" + username +
					// 	"&activekey=" + ak)
					SendRegisterMail(username, tempCheckCode)

					result := map[string]string{"status": "0", "key": sk}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					//insert fail
					result := map[string]string{"status": "5", "key": "未知原因"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}
			}
		} else {
			//it's not our client
			result := map[string]string{"status": "6", "key": "这不是我们的"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}
	} else {
		//network wrong
		result := map[string]string{"status": "4", "key": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func resendRegisterCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		fmt.Println("resend")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				tempCheckCode := getSpecificLenRandom(4)
				addHung(username, tempCheckCode)

				// SendRegisterMail(username, "http://127.0.0.1:1280/user/active/?username=" + username +
				// 	"&activekey=" + ak)
				SendRegisterMail(username, tempCheckCode)

				result := map[string]string{"status": "0", "result": "发送成功"}
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

func activeKey(username, date, random string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ran := r.Intn(90000) + 10000
	str := "username" + username + "activekey" + time.Now().String() + "wawa"
	h := md5.New()
	h.Write([]byte(str))
	a := hex.EncodeToString(h.Sum(nil))

	h1 := md5.New()
	h1.Write([]byte("random" + fmt.Sprintf("%d", ran) + "username" + username + "key" + random))
	b := hex.EncodeToString(h1.Sum(nil))
	return b + a
}

// func active(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET"{
// 		username := r.FormValue("username")
// 		activekey := r.FormValue("activekey")
// 		if userHung[username] == activekey && activekey!="" {
// 			deleteHung(username)
// 			var user User
// 			user1 := user.QueryUser(username)
// 			user1.contents["status"] = "1"

// 			if user1.update() {
// 				fmt.Fprintf(w, "激活成功！")
// 			}else {
// 				fmt.Fprintf(w, "激活失败！")
// 			}

// 		}else{
// 			fmt.Fprintf(w, "链接失效！")
// 		}
// 	}else{
// 		//network w
// 		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
// 		strResult,_ := json.Marshal(result)
// 		fmt.Fprintf(w, string(strResult))
// 	}
// }
func active(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		checkcode := r.FormValue("checkcode")
		fmt.Println("userMap:" + userMap[username].sk)
		fmt.Println("key:" + key)
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				if checkcode == userHung[username] {
					deleteHung(username)
					var user User
					user1 := user.QueryUser(username)
					user1.contents["status"] = "1"
					if user1.update() {
						result := map[string]string{"status": "0", "result": "修改成功"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					} else {
						result := map[string]string{"status": "1", "result": "修改失败"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}
				} else {
					result := map[string]string{"status": "7", "result": "　验证码错误"}
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

func checkUserid(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		var user User
		user1 := user.QueryUser(username)
		if user1.contents != nil {
			//user exist
			result := map[string]string{"status": "1", "result": "用户名存在"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		} else {
			//user not exist
			result := map[string]string{"status": "0", "result": "用户名不存在"}
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

//TODO log out and save the user state
func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				logOut(username)
				result := map[string]string{"status": "0", "result": "登出成功"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			} else {
				//key not right
				result := map[string]string{"status": "1", "result": "访问失效"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			// no login or server down
			result := map[string]string{"status": "2", "result": "已经退出"}
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

func getProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				var user User
				user1 := user.QueryUser(username)
				result := map[string]string{"status": "0",
					"username":     user1.contents["username"],
					"nickname":     user1.contents["nickname"],
					"portrait_url": user1.contents["portrait_url"],
					"gender":       user1.contents["gender"],
					"birthday":     user1.contents["birthday"],
					"state":        user1.contents["status"],
					"interest":     user1.contents["interest"],
				}
				strResult, _ := json.Marshal(result)
				fmt.Println("profile" + string(strResult))
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

func setProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		profile := r.FormValue("profile")
		fmt.Println("profile" + profile)
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				var user User
				user1 := user.QueryUser(username)
				var dat map[string]interface{}
				if err := json.Unmarshal([]byte(profile), &dat); err != nil {
					panic(err)
				}
				fmt.Println(dat)
				user1.contents["nickname"] = dat["nickname"].(string)
				user1.contents["gender"] = dat["gender"].(string)
				user1.contents["birthday"] = dat["birthday"].(string)
				// tempStr, _ := json.Marshal(dat["interest"])

				// user1.contents["interest"] =  string(tempStr)
				// fmt.Println("nickname" + user1.contents["interest"])

				if user1.update() {
					result := map[string]string{"status": "0", "result": "修改成功"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]string{"status": "5", "result": "未知错误"}
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

func uploadPortrait(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				token := image.GetUpToken(username)
				if len(token) > 0 {
					result := map[string]string{"status": "0", "token": token}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]string{"status": "5", "result": "未知错误"}
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

func setPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		key := r.FormValue("key")
		date := r.FormValue("date")
		oldpw := r.FormValue("old_password")
		newpw := r.FormValue("new_password")
		if userMap[username].sk != "" {
			if matchSessionKey(key, userMap[username].sk) {
				var user User
				user1 := user.QueryUser(username)
				temp := []byte(oldpw)
				h := md5.New()
				h.Write([]byte(user1.contents["password"] + date))
				if HexEncodeEqual(temp, h.Sum(nil)) {
					user1.contents["password"] = newpw
					fmt.Println("pass" + user1.contents["password"])
					if user1.update() {
						result := map[string]string{"status": "0", "result": "修改成功"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					} else {
						result := map[string]string{"status": "5", "result": "未知错误"}
						strResult, _ := json.Marshal(result)
						fmt.Fprintf(w, string(strResult))
					}
				} else {
					result := map[string]string{"status": "2", "result": "密码错误"}
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

// func forgetPassword(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST"{
// 		username := r.FormValue("username")
// 		date := r.FormValue("date")
// 		random := r.FormValue("random")
// 		fmt.Println(username)
// 		fmt.Println(date)
// 		fmt.Println(random)
// 		if !CheckAttack(username, date, random) {
// 			var user User
// 			user1 := user.QueryUser(username)
// 			if user1.contents !=nil{
// 				//user exist
// 				fk := forgetKey(username, date, random)
// 				addForget(username, fk)

// 				SendRegisterMail(username, mailAddress + "/user/reset_password/?username=" + username +
// 					"&forgetkey=" + fk)
// 				//user not exist
// 				result := map[string]string{"status": "0", "key": "成功"}
// 				strResult,_ := json.Marshal(result)
// 				fmt.Fprintf(w, string(strResult))
// 			}else{
// 				//user not exist
// 				result := map[string]string{"status": "2", "key": "用户名不存在"}
// 				strResult,_ := json.Marshal(result)
// 				fmt.Fprintf(w, string(strResult))
// 			}
// 		}else{
// 			//it's not our client
// 			result := map[string]string{"status": "4", "key": "这不是我们的"}
// 			strResult,_ := json.Marshal(result)
// 			fmt.Fprintf(w, string(strResult))
// 		}
// 	}else{
// 		//network wrong
// 		result := map[string]string{"status": "3", "key": "网络嗝屁了"}
// 		strResult,_ := json.Marshal(result)
// 		fmt.Fprintf(w, string(strResult))
// 	}

// }

func getVerificationCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		date := r.FormValue("date")
		random := r.FormValue("random")
		fmt.Println("date:" + date)
		fmt.Println("username:" + username)
		fmt.Println("random:" + random)
		fmt.Println("conect verfication")
		if !CheckAttack(username, date, random) {
			var user User
			user1 := user.QueryUser(username)
			if user1.contents != nil {
				//user exist
				// fk := forgetKey(username, date, random)
				checkCode := getSpecificLenRandom(6)
				addForget(username, checkCode)
				fmt.Println("conect verfication 0")
				SendForgetMail(username, checkCode)
				//user not exist
				result := map[string]string{"status": "0", "key": "成功"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			} else {
				//user not exist
				fmt.Println("conect verfication 2")
				result := map[string]string{"status": "8", "key": "用户名不存在"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			//it's not our client
			fmt.Println("conect verfication 4")
			result := map[string]string{"status": "6", "key": "这不是我们的"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}
	} else {
		//network wrong
		fmt.Println("conect verfication 3")
		result := map[string]string{"status": "4", "key": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}

}

func forgetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		date := r.FormValue("date")
		random := r.FormValue("random")
		password := r.FormValue("password")
		checkcode := r.FormValue("checkcode")

		if !CheckAttack(username, date, random) {
			if checkcode == userForget[username] {
				deleteForget(username)
				var user User
				user1 := user.QueryUser(username)
				user1.contents["password"] = password
				if user1.update() {
					sk := SessionKey(username)
					//TO-DO bugs
					addUser(username, user1.contents["id"], sk, user1.contents["interest"], date)

					result := map[string]string{"status": "0", "key": sk}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else {
					result := map[string]string{"status": "5", "result": "未知失败"}
					strResult, _ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}
			} else {
				result := map[string]string{"status": "7", "result": "验证码错误"}
				strResult, _ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		} else {
			//it's not our client
			result := map[string]string{"status": "6", "key": "这不是我们的"}
			strResult, _ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}
	} else {
		//network wrong
		result := map[string]string{"status": "4", "key": "网络嗝屁了"}
		strResult, _ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}

}

func forgetKey(username, date, random string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ran := r.Intn(90000) + 10000
	str := "username" + username + "forgetkey" + time.Now().String() + "doforget"
	h := md5.New()
	h.Write([]byte(str))
	a := hex.EncodeToString(h.Sum(nil))

	h1 := md5.New()
	h1.Write([]byte("random" + fmt.Sprintf("%d", ran) + "username" + username + "key" + random + "forget"))
	b := hex.EncodeToString(h1.Sum(nil))
	return b + a
}

// func resetPassword(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET"{
// 		username := r.FormValue("username")
// 		forgetkey := r.FormValue("forgetkey")
// 		if userForget[username] == forgetkey && forgetkey!="" {
// 			deleteForget(username)
// 			var user User
// 			user1 := user.QueryUser(username)
// 			newpwd := getNewPassword()
// 			user1.contents["password"] = newpwd
// 			if user1.update() {
// 				fmt.Fprintf(w, "新密码:"+newpwd)
// 			}else {
// 				fmt.Fprintf(w, "获取密码失败！")
// 			}

// 		}else{
// 			fmt.Fprintf(w, "链接失效！")
// 		}
// 	}else{
// 		//network w
// 		result := map[string]string{"status": "1", "result": "网络嗝屁了"}
// 		strResult,_ := json.Marshal(result)
// 		fmt.Fprintf(w, string(strResult))
// 	}
// }

const chartable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func getNewPassword(len int) string {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	var result string
// 	for i:=0; i<len; i++ {
// 		result = result + string(chartable[r.Intn(62)])
// 	}
// 	return result

// }
func getSpecificLenRandom(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result string
	for i := 0; i < len; i++ {
		result = result + string(chartable[r.Intn(62)])
	}
	return result

}

const hextable = "0123456789abcdef"

func HexEncodeEqual(dst, src []byte) bool {
	for i, v := range src {
		if dst[i*2] != hextable[v>>4] || dst[i*2+1] != hextable[v&0x0f] {
			return false
		}
	}
	return true
}

func CheckAttack(username string, date string, str string) bool {
	if username == "" || date == "" || str == "" {
		return true
	}
	h := md5.New()
	h.Write([]byte(username + "sirity"))
	a := hex.EncodeToString(h.Sum(nil))
	h1 := md5.New()
	h1.Write([]byte(date + "hello"))
	b := hex.EncodeToString(h1.Sum(nil))
	for i, _ := range a {
		if a[i] != str[i*2] || b[i] != str[i*2+1] {
			return true
		}
	}
	return false
}

func SessionKey(username string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ran := r.Intn(90000) + 10000
	str := "username" + username + "sessionkey" + time.Now().String() + "in the end"
	h := md5.New()
	h.Write([]byte(str))
	a := hex.EncodeToString(h.Sum(nil))

	h1 := md5.New()
	h1.Write([]byte("random" + fmt.Sprintf("%d", ran) + "username" + username))
	b := hex.EncodeToString(h1.Sum(nil))
	return b + a
}

func matchSessionKey(key, sk string) bool {
	fmt.Println(key)
	fmt.Println(sk)
	str := SubString(key, 0, 32)
	rs := []rune(key)
	lth := len(rs)
	str1 := "sessionkey:sirity" + sk + SubString(key, 32, lth-1) + "time"

	h := md5.New()
	h.Write([]byte(str1))
	if HexEncodeEqual([]byte(str), h.Sum(nil)) {
		return true
	} else {
		return false
	}
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendRegisterMail(to, content string) {
	user := "server@sirity.com"
	password := "dream2014"
	host := "smtp.exmail.qq.com:25"
	subject := "Active email from sirity"

	body := `
    <html>
    <body>
    <h3>
        Welcome to register our APP click this to active your account
        ` + content +
		`
    </h3>
    </body>
    </html>
    `
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}

func SendForgetMail(to, content string) {
	user := "server@sirity.com"
	password := "dream2014"
	host := "smtp.exmail.qq.com:25"
	subject := "forget password email from sirity"

	body := `
    <html>
    <body>
    <h3>
        请用下面的验证码忘记你的密码(请不要讲该验证码告诉其他人)
        ` + content +
		`
    </h3>
    </body>
    </html>
    `
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}
