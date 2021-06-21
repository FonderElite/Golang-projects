package main
import (
"fmt"
"net"
)
type Port struct{
scan_range int
}
func Scan(scope int) *Port{
  p := Port{scan_range: scope}
  for i := 1; i <= p.scan_range; i++{
go func(j int){
address := fmt.Sprintf("scanme.nmap.org:%d", j)
conn,err := net.Dial("tcp", address)
if err != nil{
return
}
conn.Close()
fmt.Printf("[+]Port %d is open!",j)
}(i)
}
return &p
}

func main(){
main_struct := Port{1024}
value := main_struct.scan_range
Scan(value)
}
