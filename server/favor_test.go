package server

import (
	"testing"
	"fmt"
)

// func TestQueryFavor(t *testing.T){
// 	Init()
// 	var favor Favor
//     comments := favor.QueryAll()
//     var le = len(comments)
//     for i := 0; i < le; i++ {
//         fmt.Println(comments[i].contents["user_id"])
//         fmt.Println(comments[i].contents["content_id"])
//         fmt.Println(comments[i].contents["date"])
//         fmt.Println("-------------------")
//     }
// }

func TestQueryFavorByUser(t *testing.T){
  Init()
  var favor Favor
    comments := favor.QueryUserId(1)
    var le = len(comments)
    for i := 0; i < le; i++ {
        fmt.Println(comments[i].contents["user_id"])
        fmt.Println(comments[i].contents["content_id"])
        fmt.Println(comments[i].contents["date"])
        fmt.Println("-------------------")
    }
}



// func TestInsertFavor(t *testing.T){
// 	Init()
// 	var favor Favor
//    	favor.Init()
//    	favor.contents["content_id"] = "1"
//    	favor.contents["user_id"] = "1"
//    	favor.contents["date"] = "2014-09-01"
//    	favor.insert()
// }



// func TestUpdateFavor(t *testing.T){
// 	Init()
//   var favor Favor
//   comments := favor.QueryUserId(1)
//   var le = len(comments)
//   for i := 0; i < le; i++ {
//       comments[i].contents["date"] = "2014-08-01"
//       comments[i].update()
//   }
// }

// func TestDeleteFavor(t *testing.T){
// 	Init()
//   var favor Favor
//   comments := favor.QueryUserId(1)
//   var le = len(comments)
//   for i := 0; i < le; i++ {
//       if comments[i].contents["content_id"] == "1"{
//         comments[i].delete()
//       }     
//   }
// }

// func BenchmarkQuery(b *testing.B) {
// 	b.StopTimer()
// 	Init()
// 	b.StartTimer()
// 	var favor Favor
// 	for i := 0; i < b.N; i++ {
// 		favor.QueryAll()
// 	}

// }
