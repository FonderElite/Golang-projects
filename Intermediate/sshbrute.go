package main
/*Source: 
https://gist.github.com/andrielfn/7aa336f9b6dc9e2a5dac
https://skarlso.github.io/2019/02/17/go-ssh-with-host-key-verification/
https://gist.github.com/montanaflynn/b59c058ce2adc18f31d6
*/
import (
  "fmt"
  "flag"
  "log"
  "time"
  "net"
  "regexp"
  "os"

)
type Verify struct{
address string
port int
}

func Banner(){
  time.Sleep(1)
  banner_print := `
  @@@@@@  @@@@@@ @@@  @@@          @@@@@@@  @@@@@@@  @@@  @@@ @@@@@@@ @@@@@@@@
 !@@     !@@     @@!  @@@          @@!  @@@ @@!  @@@ @@!  @@@   @@!   @@!     
  !@@!!   !@@!!  @!@!@!@! @!@!@!@! @!@!@!@  @!@!!@!  @!@  !@!   @!!   @!!!:!  
     !:!     !:! !!:  !!!          !!:  !!! !!: :!!  !!:  !!!   !!:   !!:     
 ::.: :  ::.: :   :   : :          :: : ::   :   : :  :.:: :     :    : :: :::
_______________________________________________________________________________

        Made By FonderElite | Github:https://github.com/FonderElite
-------------------------------------------------------------------------------
`

fmt.Println(banner_print)
time.Sleep(2)
}
func Check (host string,portnum int) *Verify{
  check_values := Verify{address: host,port: portnum}
  target := fmt.Sprintf("%s:%d", check_values.address, check_values.port)
  conn,err :=  net.DialTimeout("tcp",target,2 * time.Second)
//  derefer := &conn
  //&{{0xc00010c080}}
  r,_ := regexp.Compile("([&])?.*([0-9])(}})")
  var matchString bool = r.MatchString("&{{0x}}")
  if err != nil{
log.Fatal(err)
os.Exit(1)

}else if matchString == true{
  fmt.Println(conn)
}else{
fmt.Printf("Success %v:%v is up! \n",check_values.address,check_values.port)
//fmt.Println(conn)
defer fmt.Printf("Time: %v\n",time.Now())
  }
                return &check_values
        }
func main(){
  addrFlag := flag.String("ip","127.0.0.1","Ip address")
  portFlag := flag.Int("port",22,"Port")
  flag.Parse()
  obj := Verify{*addrFlag,*portFlag}
  var value1 = obj.address 
  var value2 = obj.port
  Banner()
Check(value1,value2)
}
