package main
import (
"fmt"
"net/http"
"flag"
"log"
//"time"
)
type Ddos struct{
domain string
port int
bytes int
} 
func Request(ipadd string) *Ddos{
ip_a := Ddos{domain: ipadd}
req,err := http.Get(ipadd)
if err != nil{
log.Fatal(err)
}else{
fmt.Println("Checking if the host is up.")
if req.StatusCode == 200 && req.StatusCode != 400{ 
fmt.Printf("[%v]Success Status-code.",req.StatusCode)
}else{
fmt.Printf("[%v]Host Seems down",req.StatusCode)
}
}
defer req.Body.Close() 
return &ip_a
}
func main(){
first_arg :=  flag.String("domain", "dns", "Domain Name")
flag.Parse()
Request(*first_arg)
}
