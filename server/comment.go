package server

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Comment struct {
	contents map[string]string
}

func (comment *Comment) Init() {
	comment.contents = make(map[string]string)
	comment.contents["id"] = ""
	comment.contents["content_id"] = ""
	comment.contents["user_id"] = ""
	comment.contents["to_comment_id"] = ""
	comment.contents["body"] = ""
	comment.contents["date"] = ""
	comment.contents["like_num"] = ""
}

func (comment Comment) QueryAll() map[int]*Comment {

	// Execute the query
	rows, err := db.Query("SELECT * FROM " + commentTable)
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

	var result map[int]*Comment
	result = make(map[int]*Comment)
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
		var comment Comment
		comment.Init()
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = ""
			} else {
				value = string(col)
			}
			// fmt.Println(columns[i], ": ", value)
			comment.contents[columns[i]] = value
		}
		// fmt.Println("-----------------------------------")
		result[index] = &comment
		index = index + 1
	}
	return result
}

func (comment Comment) QueryContentComment(contentId string) map[int]*Comment {

	// Execute the query
	rows, err := db.Query("SELECT * FROM "+commentTable+" where content_id = ? ", contentId)
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

	var result map[int]*Comment
	result = make(map[int]*Comment)
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
		var comment Comment
		comment.Init()
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = ""
			} else {
				value = string(col)
			}
			// fmt.Println(columns[i], ": ", value)
			comment.contents[columns[i]] = value
		}
		// fmt.Println("-----------------------------------")
		result[index] = &comment
		index = index + 1
	}
	return result
}

func (comment Comment) QueryId(id string) *Comment {

	// Execute the query
	rows, err := db.Query("SELECT * FROM "+commentTable+" where id = ? ", id)
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
	var temp Comment
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

func (comment *Comment) insert() bool {
	stmt, err := db.Prepare("INSERT INTO comment (id, content_id, user_id, to_comment_id, body, date, like_num )" +
		" VALUES(?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()
	checkErr(err)
	var tempDate interface{}
	if comment.contents["date"] == "" || comment.contents["date"] == "NULL" {

	} else {
		t1, err := time.Parse("2006-01-02 15:04:05", comment.contents["date"])
		if err != nil {
			fmt.Println(err)
			return false
		}
		tempDate = t1
	}

	var index interface{}
	if comment.contents["id"] == "" || comment.contents["id"] == "NULL" {

	} else {
		index = comment.contents["id"]
	}
	_, err = stmt.Exec(index, comment.contents["content_id"], comment.contents["user_id"],
		comment.contents["to_comment_id"], comment.contents["body"], tempDate, comment.contents["like_num"])
	checkErr(err)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (comment *Comment) delete() bool {
	result, err := db.Exec("DELETE FROM comment where id = ?", comment.contents["id"])
	affectedId, _ := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return false
	}
	if affectedId == 0 {
		return false
	}
	return true
}

func (comment *Comment) update() bool {
	stmt, err := db.Prepare("update comment set content_id=?, user_id=?, to_comment_id=?, body=?, date=?, like_num " +
		" where id = ?")
	checkErr(err)
	var tempDate interface{}
	if comment.contents["date"] == "" || comment.contents["date"] == "NULL" {

	} else {
		t1, err := time.Parse("2006-01-02 15:04:05", comment.contents["date"])
		if err != nil {
			fmt.Println(err)
			return false
		}
		tempDate = t1
	}

	var index interface{}
	if comment.contents["id"] == "" || comment.contents["id"] == "NULL" {

	} else {
		index = comment.contents["id"]
	}

	_, err = stmt.Exec(comment.contents["content_id"], comment.contents["user_id"],
		comment.contents["to_comment_id"], comment.contents["body"], tempDate,
		comment.contents["like_num"], index)

	checkErr(err)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
