package server

import (
	"testing"
	"fmt"
)

func TestQuery(t *testing.T){
	Init()
	var user User
    users := user.QueryAll()
    var le = len(users)
    for i := 0; i < le; i++ {
        fmt.Println(users[i].contents["id"])
        fmt.Println(users[i].contents["username"])
        fmt.Println(users[i].contents["birthday"])
        fmt.Println("-------------------")
    }
}

// func TestInsert(t *testing.T){
// 	Init()
// 	var user User
//    	user.Init()
//    	user.contents["username"] = "testinsert"
//    	user.insert()
// }

func TestQueryId(t *testing.T){
	Init()
	var user User
    user1 := user.QueryId(2)
    user1.update()
}

// func TestDelete(t *testing.T){
// 	Init()
// 	var user User
// 	user1 := user.QueryId(9)
// 	user1.delete()
// }

func BenchmarkQuery(b *testing.B) {
	b.StopTimer()
	Init()
	b.StartTimer()
	var user User
	for i := 0; i < b.N; i++ {
		user.QueryAll()
	}

}
