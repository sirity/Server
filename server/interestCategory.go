package server

import (
    "database/sql"
    "fmt"
    "log"
)

type InterestCategory struct {
    contents map[string]string
}

func (interestCategory *InterestCategory) Init() {
    interestCategory.contents = make(map[string]string)
    interestCategory.contents["id"] = ""
    interestCategory.contents["name"] = ""
    interestCategory.contents["pic"] = ""
}

func (interestCategory InterestCategory) QueryAll() map[int] *InterestCategory {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + interestCategoryTable)
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

    var result map[int] *InterestCategory
    result = make(map[int] *InterestCategory)
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
        var interestCategory InterestCategory
        interestCategory.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            interestCategory.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &interestCategory
        index = index + 1
    }
    return result
}

func (interestCategory InterestCategory) QueryInterestCategory(categoryId string) *InterestCategory {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + interestCategoryTable + " where id = ? ", categoryId)
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

    var result InterestCategory
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
        var temp InterestCategory
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

func (interestCategory *InterestCategory) insert() bool {
    stmt, err := db.Prepare("INSERT INTO interestCategory (id, name, pic)" + 
        " VALUES(?, ?, ?)")
    defer stmt.Close()
    checkErr(err)

    var index interface {}
    if interestCategory.contents["id"] == "" || interestCategory.contents["id"] == "NULL" {

    }else {
        index = interestCategory.contents["id"]
    }

    _, err = stmt.Exec(index, interestCategory.contents["name"], interestCategory.contents["pic"])
    if err != nil {
        fmt.Println(err)
        return false
    }
    checkErr(err)
    return true
}

func (interestCategory *InterestCategory) delete() bool {
    result, err := db.Exec("DELETE FROM interestCategory where id = ?", interestCategory.contents["id"])
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

func (interestCategory *InterestCategory) update() bool {
    stmt, err := db.Prepare("update interestCategory set name=?, pic=? " +
        " where id = ?")
    checkErr(err)

    var index interface {}
    if interestCategory.contents["id"] == "" || interestCategory.contents["id"] == "NULL" {

    }else {
        index = interestCategory.contents["id"]
    }
    
    _, err = stmt.Exec(interestCategory.contents["name"], interestCategory.contents["pic"], index)

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}