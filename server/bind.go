package server


func bind() {
	//user bind
    mymux.HandleFunc("/user/login", login)
    mymux.HandleFunc("/user/register", register)
    mymux.HandleFunc("/user/check_userid", checkUserid)
    mymux.HandleFunc("/user/logout", logout)
}