package server

import (
    "database/sql"
    "fmt"
    "log"
)

type InterestList struct {
    contents map[string]string
}

func (interestList *InterestList) Init() {
    interestList.contents = make(map[string]string)
    interestList.contents["id"] = ""
    interestList.contents["category"] = ""
    interestList.contents["name"] = ""
    interestList.contents["pic"] = ""
    interestList.contents["degree"] = ""
}

func (interestList InterestList) QueryAll() map[int] *InterestList {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + interestListTable)
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

    var result map[int] *InterestList
    result = make(map[int] *InterestList)
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
        var interestList InterestList
        interestList.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            interestList.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &interestList
        index = index + 1
    }
    return result
}

func (interestList InterestList) QueryInterestList(interestId string) *InterestList {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + interestListTable + " where id = ?", interestId)
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

    var result InterestList
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
        var temp InterestList
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
        result = temp
    }
    return &result
}

func (interestList *InterestList) Insert() bool {
    stmt, err := db.Prepare("INSERT INTO interestList (id, category, name, pic, degree)" + 
        " VALUES(?, ?, ?, ?, ?)")
    defer stmt.Close()
    checkErr(err)
   
    var index interface {}
    if interestList.contents["id"] == "" || interestList.contents["id"] == "NULL" {

    }else {
        index = interestList.contents["id"]
    }

    _, err = stmt.Exec(index, interestList.contents["category"], 
            interestList.contents["name"], interestList.contents["pic"], interestList.contents["degree"])
    if err != nil {
        fmt.Println(err)
        return false
    }
    checkErr(err)
    return true
}

func (interestList *InterestList) delete() bool {
    result, err := db.Exec("DELETE FROM interestList where id = ?", interestList.contents["id"])
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

func (interestList *InterestList) update() bool {
    stmt, err := db.Prepare("update interestList set category = ?, name = ?, pic = ?, degree = ? " +
        " where id = ?")
    checkErr(err)
   
    var index interface {}
    if interestList.contents["id"] == "" || interestList.contents["id"] == "NULL" {

    }else {
        index = interestList.contents["id"]
    }

    _, err = stmt.Exec(interestList.contents["category"], interestList.contents["name"], 
        interestList.contents["pic"], interestList.contents["degree"], index)

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (interestList *InterestList) SetValue(category, name, pic, degree string) {
    interestList.contents["category"] = category
    interestList.contents["name"] = name
    interestList.contents["pic"] = pic
    interestList.contents["degree"] = degree
}
