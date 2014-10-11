package server

import (
	"database/sql"
	"fmt"
	"log"
	"time"
    "strconv"
)

type Content struct {
	contents map[string]string
}

func (content *Content) Init() {
	content.contents = make(map[string]string)
	content.contents["id"] = ""
	content.contents["type"] = ""
    content.contents["mode"] = ""
	content.contents["title"] = ""
	content.contents["summary"] = ""
	content.contents["cover_url"] = ""
	content.contents["author"] = ""
	content.contents["link"] = ""
	content.contents["source"] = ""
	content.contents["content"] = ""
    content.contents["tags"] = ""
    content.contents["like_num"] = ""
    content.contents["rates"] = ""
    content.contents["rates_people"] = ""
    content.contents["date"] = ""
    content.contents["due_date"] = ""
}

func (content Content) QueryAll() map[int] *Content {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + contentTable)
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

    var result map[int] *Content
    result = make(map[int] *Content)
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
        var content Content
        content.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            content.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &content
        index = index + 1
    }
    return result
}

func (content Content) QueryRandom() map[int] *Content {

    // Execute the query
    // rows, err := db.Query("SELECT * FROM " + contentTable + "AS t1 JOIN (SELECT ROUND(RAND() * ((SELECT MAX(id) FROM `table`)-(SELECT MIN(id) FROM `table`))+(SELECT MIN(id) FROM `table`)) AS id) AS t2" + 
    //     "WHERE t1.id >= t2.id" + "ORDER BY t1.id LIMIT 7")
    rows, err := db.Query("SELECT * FROM " + contentTable + " AS t1 JOIN (SELECT ROUND(RAND() * ((SELECT MAX(id) FROM `content`)-(SELECT MIN(id) FROM `content`))+(SELECT MIN(id) FROM `content`)) AS id) AS t2 " + 
        " WHERE t1.id >= t2.id " +
        " ORDER BY t1.id LIMIT 7")
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

    var result map[int] *Content
    result = make(map[int] *Content)
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
        var content Content
        content.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            
            //remove the small id from the new content orders
            if columns[i] == "id" {
                if content.contents[columns[i]] != "" {
                    tempLast,_ := strconv.ParseInt(content.contents[columns[i]], 10, 32)
                    tempNow,_ := strconv.ParseInt(value, 10, 32)
                    if tempNow > tempLast {
                        content.contents[columns[i]] = value
                    }
                }else{
                    content.contents[columns[i]] = value
                }
            }else{
                content.contents[columns[i]] = value
            }
        }
        // fmt.Println("-----------------------------------")
        result[index] = &content
        index = index + 1
    }
    return result
}

func (content Content) QueryId(id string) *Content {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + contentTable + " where id = ? ", id)
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
    var temp Content
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
    return &temp
}


func (content *Content) Insert() bool {
    exist_con:=content.QueryAll()
   
    for _, con := range exist_con {
        if content.contents["link"] == con.contents["link"] {
           fmt.Println("link: ",con.contents["link"], "has existed")
           return false
        }
    }
    
	stmt, err := db.Prepare("INSERT INTO " + contentTable + " (id, type, mode, title, summary, cover_url, author," + 
        " link, source, content, tags, like_num, rates, rates_people, date, due_date)" + 
		" VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	checkErr(err)
    var tempDate interface{}
    if content.contents["date"] == "" ||  content.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02 15:04:05", content.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }

    var tempDueDate interface{}
    if content.contents["due_date"] == "" ||  content.contents["due_date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02 15:04:05", content.contents["due_date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDueDate = t1
    }

    var index interface {}
    if content.contents["id"] == "" || content.contents["id"] == "NULL" {

    }else {
        index = content.contents["id"]
    }

	_, err = stmt.Exec(index, content.contents["type"], content.contents["mode"], content.contents["title"], 
        content.contents["summary"], content.contents["cover_url"], content.contents["author"], content.contents["link"],
        content.contents["source"], content.contents["content"], content.contents["tags"], content.contents["like_num"],
        content.contents["rates"], content.contents["rates_people"], tempDate, tempDueDate,)
	checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (content *Content) delete(){
	_, err := db.Exec("DELETE FROM content where id = ?", content.contents["id"])
	if err != nil {
		log.Println(err)
		return
	}
}

func (content *Content) update() bool {
	stmt, err := db.Prepare("update content set type=?, mode=?, title=?, summary=?, cover_url=?, author=?, " +
		"link=?, source=?, content=?, tags=?, like_num=?, rates=?, rates_people=?, date=?, due_date=? where id = ?")
    checkErr(err)

    var tempDate interface{}
    if content.contents["date"] == "" ||  content.contents["date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02 15:04:05", content.contents["date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDate = t1
    }

    var tempDueDate interface{}
    if content.contents["due_date"] == "" ||  content.contents["due_date"] == "NULL"{

    }else{
        t1, err := time.Parse("2006-01-02 15:04:05", content.contents["due_date"])
        if err != nil {
            fmt.Println(err)
            return false
        }
        tempDueDate = t1
    }

    var index interface {}
    if content.contents["id"] == "" || content.contents["id"] == "NULL" {

    }else {
        index = content.contents["id"]
    }

    _, err = stmt.Exec(content.contents["type"], content.contents["mode"], content.contents["title"], content.contents["summary"],
        content.contents["cover_url"], content.contents["author"], content.contents["link"], content.contents["source"],
        content.contents["content"], content.contents["tags"], content.contents["like_num"], content.contents["rates"],
        content.contents["rates_people"], tempDate, tempDueDate, index)

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (content *Content) SetValue(tempType, mode, title, summary, coverUrl,
     author, link, source, tempContent, tags, likeNum, rates, ratesPeople, tempDate, dueDate string) {
    content.contents["type"] = tempType
    content.contents["mode"] = mode
    fmt.Println("mode: ",mode)
    content.contents["title"] = title
    content.contents["summary"] = summary
    content.contents["cover_url"] = coverUrl
    content.contents["author"] = author
    content.contents["link"] = link
    content.contents["source"] = source
    content.contents["content"] = tempContent
    content.contents["tags"] = tags
    content.contents["like_num"] = likeNum
    content.contents["rates"] = rates
    content.contents["rates_people"] = ratesPeople
    content.contents["date"] = tempDate
    content.contents["due_date"] = dueDate
}


