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
    resp, err := http.PostForm("http://127.0.0.1:1280/user/login",
        url.Values{"username": {"test阿"}, "password": {"a670925a0a51e179f1343e8deb46dff7"}, "date": {"2014-8-10"}, "random": {"1299"}})
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
