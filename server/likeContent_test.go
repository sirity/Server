package server

import (
	"testing"
	"fmt"
)

// func TestQueryLikeContent(t *testing.T){
// 	Init()
// 	var likeContent LikeContent
//     comments := likeContent.QueryAll()
//     var le = len(comments)
//     for i := 0; i < le; i++ {
//         fmt.Println(comments[i].contents["user_id"])
//         fmt.Println(comments[i].contents["content_id"])
//         fmt.Println(comments[i].contents["date"])
//         fmt.Println("-------------------")
//     }
// }

func TestQueryLikeContentByUser(t *testing.T){
  Init()
  var likeContent LikeContent
    comments := likeContent.QueryUserId(1)
    var le = len(comments)
    for i := 0; i < le; i++ {
        fmt.Println(comments[i].contents["user_id"])
        fmt.Println(comments[i].contents["content_id"])
        fmt.Println(comments[i].contents["date"])
        fmt.Println("likecontent-------------------")
    }
}



// func TestInsertLikeContent(t *testing.T){
// 	Init()
// 	var likeContent LikeContent
//    	likeContent.Init()
//    	likeContent.contents["content_id"] = "1"
//    	likeContent.contents["user_id"] = "1"
//    	likeContent.contents["date"] = "2014-09-01"
//    	likeContent.insert()
// }



// func TestUpdateLikeContent(t *testing.T){
// 	Init()
//   var likeContent LikeContent
//   comments := likeContent.QueryUserId(1)
//   var le = len(comments)
//   for i := 0; i < le; i++ {
//       comments[i].contents["date"] = "2014-08-01"
//       comments[i].update()
//   }
// }

// func TestDeleteLikeContent(t *testing.T){
// 	Init()
//   var likeContent LikeContent
//   comments := likeContent.QueryUserId(1)
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
// 	var likeContent LikeContent
// 	for i := 0; i < b.N; i++ {
// 		likeContent.QueryAll()
// 	}

// }
