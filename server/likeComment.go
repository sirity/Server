package server

import (
    "database/sql"
    "fmt"
    "log"
    "time"
)

type LikeComment struct {
    contents map[string]string
}

func (likeComment *LikeComment) Init() {
    likeComment.contents = make(map[string]string)
    likeComment.contents["user_id"] = ""
    likeComment.contents["comment_id"] = ""
    likeComment.contents["date"] = ""
}

func (likeComment LikeComment) QueryAll() map[int] *LikeComment {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + likeCommentTable)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    var result map[int] *LikeComment
    result = make(map[int] *LikeComment)
    var index int
    index = 0
    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        var likeComment LikeComment
        likeComment.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            likeComment.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &likeComment
        index = index + 1
    }
    return result
}

func (likeComment LikeComment) QueryUserId(userId int) map[int] *LikeComment {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + likeCommentTable + " where user_id = ? ", userId)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    var result map[int] *LikeComment
    result = make(map[int] *LikeComment)
    index := 0
    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        var temp LikeComment
        temp.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            temp.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &temp
        index = index + 1
    }
    return result
}

func (likeComment *LikeComment) insert() bool {
    stmt, err := db.Prepare("INSERT INTO likeComment (user_id, comment_id, date)" + 
        " VALUES(?, ?, ?)")
    defer stmt.Close()
    checkErr(err)
    var tempDate interface{}
    if likeComment.contents["date"] == "" ||  likeComment.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02", likeComment.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }

    _, err = stmt.Exec(likeComment.contents["user_id"], likeComment.contents["comment_id"], tempDate)
    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (likeComment *LikeComment) delete(){
    _, err := db.Exec("DELETE FROM likeComment where user_id = ? & comment_id = ?", likeComment.contents["user_id"], likeComment.contents["comment_id"])
    if err != nil {
        log.Println(err)
        return
    }
}

func (likeComment *LikeComment) update() bool {
    stmt, err := db.Prepare("update likeComment set date=? " +
        " where user_id = ? & comment_id=?")
    checkErr(err)
    var tempDate interface{}
    if likeComment.contents["date"] == "" ||  likeComment.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02", likeComment.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }
    
    _, err = stmt.Exec(tempDate, likeComment.contents["user_id"], likeComment.contents["comment_id"])

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}