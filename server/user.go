package server

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type User struct {
	contents map[string]string
}

func (user *User) Init() {
	user.contents = make(map[string]string)
	user.contents["id"] = ""
	user.contents["username"] = ""
	user.contents["nickname"] = ""
	user.contents["password"] = ""
	user.contents["portraitUrl"] = ""
	user.contents["gender"] = ""
	user.contents["birthday"] = ""
	user.contents["status"] = ""
	user.contents["interest"] = ""
    user.contents["category"] = ""
}

func (user User) QueryAll() map[int] *User {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + userTable)
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

    var result map[int] *User
    result = make(map[int] *User)
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
        var user User
        user.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            user.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &user
        index = index + 1
    }
    return result
}

func (user User) QueryId(id string) User {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + userTable + " where id = ? ", id)
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
    var temp User
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
        index = index + 1
    }
    return temp
}

func (user User) QueryUser(username string) User {

    // Execute the query
    if db == nil {
        fmt.Println("db is null")
    }
    rows, err := db.Query("SELECT * FROM " + userTable + " where username = ? ", username)
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
    var temp User
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
        index = index + 1
    }
    return temp
}

func (user *User) insert() bool {
	stmt, err := db.Prepare("INSERT INTO user (id, username, nickname, password, portraitUrl, gender, birthday, status, interest, category)" + 
		" VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	checkErr(err)
    var t interface{}
    if user.contents["birthday"] == "" ||  user.contents["birthday"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02", user.contents["birthday"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        t = t1
    }

    var index interface {}
    if user.contents["id"] == "" || user.contents["id"] == "NULL" {

    }else {
        index = user.contents["id"]
    }
	_, err = stmt.Exec(index, user.contents["username"], user.contents["nickname"],
		user.contents["password"], user.contents["portraitUrl"], user.contents["gender"],
		t, user.contents["status"], user.contents["interest"], user.contents["category"])
	checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (user *User) delete(){
	_, err := db.Exec("DELETE FROM user where id = ?", user.contents["id"])
	if err != nil {
		log.Println(err)
		return
	}
}

func (user *User) update() bool {
	stmt, err := db.Prepare("update user set username=?, nickname=?, password=?, portraitUrl=?, gender=?, " +
		"birthday=?, status=?, interest=?, category=? where id = ?")
    checkErr(err)
    var t interface{}
    if user.contents["birthday"] == "" ||  user.contents["birthday"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02", user.contents["birthday"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        t = t1
    }

    var index interface {}
    if user.contents["id"] == "" || user.contents["id"] == "NULL" {

    }else {
        index = user.contents["id"]
    }
    
    _, err = stmt.Exec(user.contents["username"], user.contents["nickname"],
        user.contents["password"], user.contents["portraitUrl"], user.contents["gender"],
        t, user.contents["status"], user.contents["interest"], user.contents["category"], index)

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}


