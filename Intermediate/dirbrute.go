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
list,err := os.Open(req.wordlist)
if err != nil{
log.Fatal(err)
}
defer list.Close()
scanner := bufio.NewScanner(list)
scanner.Split(bufio.ScanWords)
for scanner.Scan() {
        wlist := scanner.Text()
for _,line := range wlist{
        request,err := http.Get(string(req.url) + "/" + string(line))
       if err != nil{
       log.Fatal(err)
       }
       fmt.Printf("[%v]Request-> /%v\n",request.StatusCode,line)
    }
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
