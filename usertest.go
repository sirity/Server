package main  
  
import (  
    "fmt"  
    "net/http"
    "io/ioutil"
    "strings"
)  
  
func main() {  
    httpPost()
} 


func httpPost() {
    resp, err := http.Post("http://127.0.0.1:1280/user/login",
        "application/x-www-form-urlencoded",
        strings.NewReader("username=ben"))
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
