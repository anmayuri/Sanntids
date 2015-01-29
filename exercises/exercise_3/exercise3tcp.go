// go 1.2 
// go run exercise3tcp.go 

package main 

import(
	"fmt"
	"net"
	"time"
	"strings"
)


func handleRequestTCP(connection net.Conn){
	tcpsend := []byte("send this tcp\x00")
	tcpreceivebuf := make([]byte, 1024)
	connection.Write(tcpsend)
	connection.Read(tcpreceivebuf)
	fmt.Println(string(tcpreceivebuf))
	time.Sleep(1*time.Second)
}

func main(){
	tcpcon, err3 := net.Dial("tcp", "localhost"+"32933")
	if err3 != nil {
		fmt.Println("error 3: tcp connection", err3)
		return
	}
	tcpbuf := make([]byte, 1024)
	fmt.Println("success")
	
	go handleRequestTCP(newconnection)
	return
}
