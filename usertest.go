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
        url.Values{"username": {"admin"}, "password": {"a670925a0a51e179f1343e8deb46dff7"}, "date": {"1408418639.363336"}, "random": {"4239caa8481680fa2d65b11f415c741ac9b78f49fda00ac466fae94e2f00604a"}})
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

func httpPost1() {
    resp, err := http.PostForm("http://121.40.94.51:1280/plan/fetch_plan",
        url.Values{"username": {"admin"}, "password": {"a670925a0a51e179f1343e8deb46dff7"}, "date": {"1408418639.363336"}, "random": {"4239caa8481680fa2d65b11f415c741ac9b78f49fda00ac466fae94e2f00604a"}})
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
