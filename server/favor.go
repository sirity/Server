package server

import (
    "database/sql"
    "fmt"
    "log"
    "time"
)

type Favor struct {
    contents map[string]string
}

func (favor *Favor) Init() {
    favor.contents = make(map[string]string)
    favor.contents["user_id"] = ""
    favor.contents["content_id"] = ""
    favor.contents["date"] = ""
}

func (favor Favor) QueryAll() map[int] *Favor {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + favorTable)
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

    var result map[int] *Favor
    result = make(map[int] *Favor)
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
        var favor Favor
        favor.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            favor.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &favor
        index = index + 1
    }
    return result
}

func (favor Favor) QueryUserId(userId int) map[int] *Favor {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + favorTable + " where user_id = ? ", userId)
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

    var result map[int] *Favor
    result = make(map[int] *Favor)
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
        var temp Favor
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

func (favor *Favor) insert() bool {
    stmt, err := db.Prepare("INSERT INTO favor (user_id, content_id, date)" + 
        " VALUES(?, ?, ?)")
    defer stmt.Close()
    checkErr(err)
    var tempDate interface{}
    if favor.contents["date"] == "" ||  favor.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02", favor.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }

    _, err = stmt.Exec(favor.contents["user_id"], favor.contents["content_id"], tempDate)
    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (favor *Favor) delete(){
    _, err := db.Exec("DELETE FROM favor where user_id = ? & content_id = ?", favor.contents["user_id"], favor.contents["content_id"])
    if err != nil {
        log.Println(err)
        return
    }
}

func (favor *Favor) update() bool {
    stmt, err := db.Prepare("update favor set date=? " +
        " where user_id = ? & content_id=?")
    checkErr(err)
    var tempDate interface{}
    if favor.contents["date"] == "" ||  favor.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02", favor.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }
    
    _, err = stmt.Exec(tempDate, favor.contents["user_id"], favor.contents["content_id"])

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}