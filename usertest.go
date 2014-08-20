package main  
  
import (  
    "fmt"  
    "net/http"
    "net/url" 
    "io/ioutil"
)  
  
func main() {  
    httpPostRegister()
    // httpPostCheckId()
} 


func httpPostLogin() {
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

func httpPostFetchPlan() {
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

func httpPostRegister() {
    resp, err := http.PostForm("http://127.0.0.1:1280/user/register",
        url.Values{"username": {"printf_ll@qq.com"}, "password": {"a670925a0a51e179f1343e8deb46dff7"}, "date": {"1408501471.322662"}, 
        "random": {"54867235a2215cf0f1ddcef8c42cfc04d005a07b5993f3d9ba963728c771dbfd"}})
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

func httpPostCheckId() {
    resp, err := http.PostForm("http://192.168.1.126:1280/user/check_userid",
        url.Values{"username": {"printfldl@gmail.com"}, "password": {"a670925a0a51e179f1343e8deb46dff7"}, "date": {"1408436324.152133"}, 
        "random": {"b9400c6b955e86e3012f07a6bb06dd0a866f8f31b3bc773c9f1912a9c588fed7"}})
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
