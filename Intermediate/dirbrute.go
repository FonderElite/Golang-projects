package main
import (
        "fmt"
        "net/http"
        "log"
        "os"
        "bufio"
)

type Request struct{
url string
wordlist string
}
func BruteForce(url string, wordl string) *Request{
req := Request{url:url,wordlist: wordl}
req.url = url
resp, err := http.Get(req.url)
if err != nil{
log.Fatal(err)
}else if resp.StatusCode == 200{
fmt.Printf("[%v]Success Status Code.",resp.StatusCode)
openf,err := os.Open(req.wordlist)
if err != nil{
log.Fatal(err)
}
defer openf.Close()
scanner := bufio.NewScanner(openf)
scanner.Split(bufio.ScanLines)
var fileTextLines []string
for scanner.Scan() {
                fileTextLines = append(fileTextLines, scanner.Text())
        }
openf.Close()
for _, eachline := range fileTextLines {
request,err := http.Get(string(req.url) + "/" + string(eachline))
if err != nil {
log.Fatal(err)
}
fmt.Printf("[%v]Request-> /%v \n",request.StatusCode,eachline)
}
}else{
fmt.Println("Error.")
}
return &req
}
func main(){
var site string
var wordlist string
fmt.Println("-HTTP REQUEST-By FonderElite-")
fmt.Print("Enter Url: ")
fmt.Scanln(&site)
fmt.Print("Wordlist: ")
fmt.Scanln(&wordlist)
BruteForce(site,wordlist)
}
