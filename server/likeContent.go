package server

import (
    "database/sql"
    "fmt"
    "log"
    "time"
)

type LikeContent struct {
    contents map[string]string
}

func (likeContent *LikeContent) Init() {
    likeContent.contents = make(map[string]string)
    likeContent.contents["user_id"] = ""
    likeContent.contents["content_id"] = ""
    likeContent.contents["date"] = ""
}

func (likeContent LikeContent) QueryAll() map[int] *LikeContent {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + likeContentTable)
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

    var result map[int] *LikeContent
    result = make(map[int] *LikeContent)
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
        var likeContent LikeContent
        likeContent.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            likeContent.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &likeContent
        index = index + 1
    }
    return result
}

func (likeContent LikeContent) QueryUserId(userId int) map[int] *LikeContent {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + likeContentTable + " where user_id = ? ", userId)
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

    var result map[int] *LikeContent
    result = make(map[int] *LikeContent)
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
        var temp LikeContent
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

func (likeContent LikeContent) QueryLikeContent(userId, contentId string) *LikeContent {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + likeContentTable + " where user_id = ? AND content_id = ?", userId, contentId)
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

    var result *LikeContent
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
        var temp LikeContent
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
        result = &temp
    }
    return result
}

func (likeContent *LikeContent) insert() bool {
    stmt, err := db.Prepare("INSERT INTO likeContent (user_id, content_id, date)" + 
        " VALUES(?, ?, ?)")
    defer stmt.Close()
    checkErr(err)
    var tempDate interface{}
    if likeContent.contents["date"] == "" ||  likeContent.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02 15:04:05", likeContent.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }

    _, err = stmt.Exec(likeContent.contents["user_id"], likeContent.contents["content_id"], tempDate)
    if err != nil {
        fmt.Println(err)
        return false
    }
    checkErr(err)
    return true
}

func (likeContent *LikeContent) delete() bool {
    result, err := db.Exec("DELETE FROM likeContent where user_id = ? AND content_id = ?", likeContent.contents["user_id"], likeContent.contents["content_id"])
    affectedId,_ := result.RowsAffected()
    if err != nil {
        log.Println(err)
        return false
    }
    if affectedId == 0 {
        return false
    }
    return true
}

func (likeContent *LikeContent) update() bool {
    stmt, err := db.Prepare("update likeContent set date=? " +
        " where user_id = ? AND content_id=?")
    checkErr(err)
    var tempDate interface{}
    if likeContent.contents["date"] == "" ||  likeContent.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02 15:04:05", likeContent.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }
    
    _, err = stmt.Exec(tempDate, likeContent.contents["user_id"], likeContent.contents["content_id"])

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}