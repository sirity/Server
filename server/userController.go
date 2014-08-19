package server

import (
	"net/http"
	"fmt"
	"encoding/json"
	"crypto/md5"
    "encoding/hex"
    "time"
    "math/rand"
    "strings"
    "net/smtp"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		password := r.FormValue("password")
		date := r.FormValue("date")
		random := r.FormValue("random")
		if !CheckAttack(username, date, random) {
			var user User
			user1 := user.QueryUser(username)
			if user1.contents !=nil{
				//user exist
				temp := []byte(password)
			    h := md5.New()
			    h.Write([]byte(user1.contents["password"] + date))
				if HexEncodeEqual(temp, h.Sum(nil)) {
					//login success
					sk := SessionKey(username)
					addUser(username, sk, user1.contents["interest"], date)
					result := map[string]string{"status": "0", "key": sk}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				} else{
					//password fail
					result := map[string]string{"status": "1", "key": "密码错误"}
					strResult,_ := json.Marshal(result)
					fmt.Fprintf(w, string(strResult))
				}
			}else{
				//user not exist
				result := map[string]string{"status": "2", "key": "用户名不存在"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			//it's not our client
			result := map[string]string{"status": "4", "key": "这不是我们的"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}
	}else{
		//network wrong
		result := map[string]string{"status": "3", "key": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
	
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		password := r.FormValue("password")
		date := r.FormValue("date")
		random := r.FormValue("random")
		if !CheckAttack(username, date, random) {
			var user User
			user1 := user.QueryUser(username)
			if user1.contents !=nil{
				//user exist
				result := map[string]string{"status": "1", "result": "用户名存在"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}else{
				//user not exist
				var user User 
				user.Init()
				user.contents["username"] = username
				user.contents["nickname"] = "我是小白"
				user.contents["password"] = password
				user.contents["status"] = "0"
				user.insert()

				// SendRegisterMail(username, content)
				result := map[string]string{"status": "0", "result": "注册成功"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			//it's not our client
			result := map[string]string{"status": "4", "key": "这不是我们的"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}
	}else{
		//network wrong
		result := map[string]string{"status": "3", "key": "网络嗝屁了"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}


func checkUserid(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		var user User
		user1 := user.QueryUser(username)
		if user1.contents !=nil{
			//user exist
			result := map[string]string{"status": "1", "result": "用户名存在"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}else{
			//user not exist
			result := map[string]string{"status": "0", "result": "用户名不存在"}
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

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "loginout")
}

func getUserInfo(w http.ResponseWriter, r *http.Request) {
	
}

func setUserInfo(w http.ResponseWriter, r *http.Request) {
	
}

func uploadPortrait(w http.ResponseWriter, r *http.Request) {

}

func setPassword(w http.ResponseWriter, r *http.Request) {

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
	h := md5.New()
	h.Write([]byte(username + "sirity"))
	a := hex.EncodeToString(h.Sum(nil))
	h1 := md5.New()
	h1.Write([]byte(date + "hello"))
	b := hex.EncodeToString(h1.Sum(nil))
	fmt.Println(a);
	fmt.Println(b);
	for i, _ := range a {
		if a[i] != str[i*2] || b[i] != str[i*2 + 1] {
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
	h1.Write([]byte("random" + fmt.Sprintf("%d", ran)  + "username" + username))
	b := hex.EncodeToString(h1.Sum(nil))
	return b + a
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
func SendMail(user, password, host, to, subject, body, mailtype string) error{
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/"+ mailtype + "; charset=UTF-8"
    }else{
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }

    msg := []byte("To: " + to + "\r\nFrom: " + user + "<"+ user +">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    return err
}

func SendRegisterMail(to, content string){
    user := "server@sirity.com"
    password := "dream2014"
    host := "smtp.exmail.qq.com:25"
    subject := "Active email from sirity"

    body := `
    <html>
    <body>
    <h3>
        Welcome to register our APP click this to active your account
        `+ content +
        `
    </h3>
    </body>
    </html>
    `
    err := SendMail(user, password, host, to, subject, body, "html")
    if err != nil {
        fmt.Println("send mail error!")
        fmt.Println(err)
    }else{
        fmt.Println("send mail success!")
    }
}



