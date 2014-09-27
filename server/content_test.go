package server

import (
	"testing"
	"fmt"
)

func TestQueryContent(t *testing.T){
	Init()
	var content Content
    users := content.QueryAll()
    var le = len(users)
    for i := 0; i < le; i++ {
        fmt.Println(users[i].contents["id"])
        fmt.Println(users[i].contents["type"])
        fmt.Println(users[i].contents["title"])
        fmt.Println("-------------------")
    }
}

// func TestInsertContent(t *testing.T){
// 	Init()
// 	var content Content
//    	content.Init()
//    	content.contents["type"] = "3"
//    	content.contents["title"] = "testinsert"
//    	content.insert()
// }

// func TestUpdateContent(t *testing.T){
// 	Init()
// 	var content Content
// 	content1 := content.QueryId(2)
// 	content1.contents["author"] = "hari"
// 	content1.update()
// }



// func TestQueryId(t *testing.T){
// 	Init()
// 	var content Content
//     user1 := content.QueryId(2)
//     user1.update()
// }

// func TestDelete(t *testing.T){
// 	Init()
// 	var content Content
// 	user1 := content.QueryId(9)
// 	user1.delete()
// }

// func BenchmarkQuery(b *testing.B) {
// 	b.StopTimer()
// 	Init()
// 	b.StartTimer()
// 	var content Content
// 	for i := 0; i < b.N; i++ {
// 		content.QueryAll()
// 	}

// }
