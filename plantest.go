package main  
  
import (  
    "fmt"  
    "net/http"
    "net/url" 
    "io/ioutil"
)  
  
// func main() {  
//     // httpPostToggleFavor()
//     httpPostToggleLikeCotnent()
//     // httpPostFetchComment()
//     // httpPostCheckId()
// }

func httpPostToggleFavor() {
    resp, err := http.PostForm("http://127.0.0.1:1280/favor/toggle_favor",
        url.Values{"username": {"printfldl@gmail.com"}, "key": {"a670925a0a51e179f1343e8deb46dff7"}, "content_id": {"2"}, 
            "last_status":{"0"}})
    if err != nil {
        fmt.Println(err)
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
    fmt.Println("tail")

}

func httpPostToggleLikeCotnent() {
    resp, err := http.PostForm("http://127.0.0.1:1280/content/toggle_like",
        url.Values{"username": {"printfldl@gmail.com"}, "key": {"a670925a0a51e179f1343e8deb46dff7"}, "content_id": {"2"}, 
            "last_status":{"1"}})
    if err != nil {
        fmt.Println(err)
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
    fmt.Println("tail")

}
