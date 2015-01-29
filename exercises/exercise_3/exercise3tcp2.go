// go 1.2 
// go run exercise3tcp2.go 

package main 

import(
	"fmt"
	"net"
	//"time"
	//"strings"
)

const (
	LOCAL_HOST = ""
	LOCAL_PORT = ":34933"
	CONN_TYPE = "tcp"
	B_PORT = ":30000"
) 

func main(){
	//lAddr, _ := net.ResolveTCPAddr(CONN_TYPE,B_PORT)
	l1, err1 := net.Listen(CONN_TYPE,LOCAL_PORT)
	if err1 != nil {
		fmt.Println("error listening:", err1)
		return
	}
	for{
		connection, err2 := l1.Accept()
		if err2 != nil {
			fmt.Println("error accepting: ", err2)
			return
		}
		go handleRequestTCP(connection)
	}
	
}

func handleRequestTCP(connection net.Conn){
	buffer := make([]byte, 1024)
	_, err3 := connection.Read(buffer)
	if err3 != nil {
		fmt.Println("error reading:", err3)
		return
	}
	message := []byte("424242\x00")
	connection.Write(message)
}


