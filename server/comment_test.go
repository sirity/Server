package server

import (
	"testing"
	"fmt"
)

func TestQueryComment(t *testing.T){
	Init()
	var comment Comment
    comments := comment.QueryAll()
    var le = len(comments)
    for i := 0; i < le; i++ {
        fmt.Println(comments[i].contents["id"])
        fmt.Println(comments[i].contents["to_comment_id"])
        fmt.Println(comments[i].contents["body"])
        fmt.Println("-------------------")
    }
}

// func TestInsertComment(t *testing.T){
// 	Init()
// 	var comment Comment
//    	comment.Init()
//    	comment.contents["content_id"] = "2"
//    	comment.contents["user_id"] = "2"
//    	comment.contents["body"] = "testinsert"
//    	comment.insert()
// }

// func TestQueryId(t *testing.T){
// 	Init()
// 	var comment Comment
//     user1 := comment.QueryId(2)
//     user1.update()
// }

// func TestDelete(t *testing.T){
// 	Init()
// 	var comment Comment
// 	user1 := comment.QueryId(9)
// 	user1.delete()
// }

// func BenchmarkQuery(b *testing.B) {
// 	b.StopTimer()
// 	Init()
// 	b.StartTimer()
// 	var comment Comment
// 	for i := 0; i < b.N; i++ {
// 		comment.QueryAll()
// 	}

// }
