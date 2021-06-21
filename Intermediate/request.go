package main
import (
        "fmt"
        "net/http"
        "log"
)

type Request struct{
url string
wordlist string
}
func CheckSite(url string) *Request{
req = Request{url:url}
req.url = url
resp, err := http.Get(req.url)
if err != nil{
log.Fatal(err)
}else if resp.statusCode == 200{
fmt.Printf("[%v]Success Status Code.",resp.statusCode)
        fmt.Println(resp.Body)
}else{
fmt.Println("Error.")
}
defer resp.Body.Close()
return &req
}
func main(){
        fmt.Println(
"     ___ ___  __      __   ___  __        ___  __  ___ "
"|__|  |   |  |__)    |__) |__  /  \ |  | |__  /__`  |  "
"|  |  |   |  |       |  \ |___ \__X \__/ |___ .__/  |  "
"By FonderElite | Github:https://github.com/FonderElite")
var site string
fmt.Println("Enter Url: ")
fmt.Scanln(&site)
CheckSite(site)
}
