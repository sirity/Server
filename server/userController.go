package server

import (
	"net/http"
	"fmt"
	"encoding/json"
)


func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		password := r.FormValue("password")
		date := r.FormValue("date")
		random := r.FormValue("random")


		result := map[string]string{"status": "0", "key": username + password + date + random}
		strResult,_ := json.Marshal(result);
		fmt.Fprintf(w, string(strResult));
	}
	
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		username := r.FormValue("username")
		password := r.FormValue("password")
		date := r.FormValue("date")
		random := r.FormValue("random")


		result := map[string]string{"status": "0", "key": "jsikeoadfoutdjjfooo"}
		strResult,_ := json.Marshal(result);
		fmt.Fprintf(w, string(strResult));
	}
	fmt.Fprintf(w, "register")
}

func checkUserid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "checkUserid")
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