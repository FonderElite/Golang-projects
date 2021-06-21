package main
import (
        "fmt"
        "net/http"
        "log"
        "io/ioutil"
)

type Request struct{
url string
wordlist string
}
func CheckSite(url string) *Request{
req := Request{url:url}
req.url = url
resp, err := http.Get(req.url)
defer resp.Body.Close()
body,err := ioutil.ReadAll(resp.Body)
if err != nil{
log.Fatal(err)
}else if resp.StatusCode == 200{
fmt.Printf("[%v]Success Status Code.",resp.StatusCode)
fmt.Printf("Content: %v", string(body))
}else{
fmt.Println("Error.")
}
return &req
}
func main(){
var site string
fmt.Println("-HTTP REQUEST-By FonderElite-")
fmt.Print("Enter Url: ")
fmt.Scanln(&site)
CheckSite(site)
}
