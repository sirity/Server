package main  
  
import (  
    "fmt"  
    "net/http"
    "net/url" 
    "io/ioutil"
)  
  
func main() {  
    httpPost()
} 


func httpPost() {
    resp, err := http.PostForm("http://121.40.94.51:1280/user/login",
        url.Values{"username": {"Value"}, "password": {"123"}, "date": {"2014-8-10"}, "random": {"1299"}})
    if err != nil {
        fmt.Println(err)
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 
    fmt.Println(string(body))
}
