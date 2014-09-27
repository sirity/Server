package server


func bind() {
	//user bind
    mymux.HandleFunc("/user/login", login)
    mymux.HandleFunc("/user/register", register)
    mymux.HandleFunc("/user/active/", active)
    mymux.HandleFunc("/user/check_userid", checkUserid)
    mymux.HandleFunc("/user/logout", logout)
    mymux.HandleFunc("/user/forget_password", forgetPassword)
    mymux.HandleFunc("/user/set_password", setPassword)
    mymux.HandleFunc("/user/set_profile", setProfile)
    mymux.HandleFunc("/user/get_profile", getProfile)
    mymux.HandleFunc("/user/reset_password/", resetPassword)



    //plan bind
    mymux.HandleFunc("/plan/fetch_plan", fetchPlan)
    mymux.HandleFunc("/plan/fetch_content_item", fetchContentItem)
    mymux.HandleFunc("/favor/toggle_favor", toggleFavor)
    mymux.HandleFunc("/favor/fetch_favor_list", fetchFavorList)
    mymux.HandleFunc("/content/toggle_like", toggleLike)
    mymux.HandleFunc("/content/fetch_like_number", fetchLikeNumber)
    mymux.HandleFunc("/content/do_i_like", doILike)


    //comment bind
    mymux.HandleFunc("/comment/fetch_comment_list", fetchCommentList)
    mymux.HandleFunc("/comment/post_comment", postComment)
    mymux.HandleFunc("/comment/delete_comment", deleteComment)
    mymux.HandleFunc("/comment/toggle_like", commentToggleLike)
    mymux.HandleFunc("/comment/fetch_like_number", commentFetchLikeNumber)
    mymux.HandleFunc("/comment/do_i_like", commentDoILike)
}